package milvus

import (
	"context"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/klimenkoOleg/migrate-stock-data/internal/app"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"go.uber.org/zap"
	"log"
	"math"
	"strconv"
	"time"
)

const (
	collectionName = "ticks1day"
)

type Milvus struct {
	config *app.Config
	log    *zap.Logger
	client client.Client
	data   map[int64][]float64
	dates  map[int64][]string
}

func NewMilvus(config *app.Config, log *zap.Logger) *Milvus {
	return &Milvus{
		config: config,
		log:    log,
		data:   make(map[int64][]float64),
		dates:  make(map[int64][]string),
	}
}

func (m *Milvus) prepare(ctx context.Context, dim int64) {
	c, err := client.NewClient(ctx, client.Config{
		Address: m.config.MilvusServer,
	})

	if err != nil {
		// handling error and exit, to make example simple here
		log.Fatal("failed to connect to milvus:", err.Error())
	}
	// in a main func, remember to close the client

	has, err := c.HasCollection(ctx, collectionName)
	if err != nil {
		log.Fatal("failed to check whether collection exists:", err.Error())
	}
	if has {
		// collection with same name exist, clean up mess
		if err := c.DropCollection(ctx, collectionName); err != nil {
			m.log.Error("failed to drop collection:", zap.Error(err))
		}
	}

	// define collection schema, see tick1day.csv
	schema := entity.NewSchema().WithName(collectionName).WithDescription("`Stock exchange daily ticks`").
		WithField(entity.NewField().WithName("ID").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true)).
		//WithField(entity.NewField().WithName("StockId").WithDataType(entity.FieldTypeInt64)).
		WithField(entity.NewField().WithName("Vector").WithDataType(entity.FieldTypeFloatVector).WithDim(dim))

	err = c.CreateCollection(ctx, schema, entity.DefaultShardNumber) // only 1 shard
	if err != nil {
		log.Fatal("failed to create collection:", err.Error())
	}
	//if len(m.ids) == 0 {
	//	m.ids = make([]int64, 0, len(ticks))
	//m.vectors = make([][]float32, 0, len(ticks))
	//}

	m.client = c
	//m.data = make(map[int64][]float32)
}

func (m *Milvus) WriteTick(ticks []dto.Tick1Day) error {
	for _, tick := range ticks {
		vector, ok := m.data[tick.InstrumentId]
		if !ok {
			diff := m.config.DateEnded.Sub(*m.config.DateStarted)
			days := int64(diff.Hours() / 24)
			vector = make([]float64, days)
		}
		id := dateToId(tick.Date, *m.config.DateStarted)
		vector[id] = float64(tick.Avg / 100_000_000)
		m.data[tick.InstrumentId] = vector
	}

	return nil
}

func (m *Milvus) Flush(ctx context.Context) error {
	ids := make([]int64, 0, len(m.data))
	vectors := make([][]float64, 0, len(m.data))

	for k, v := range m.data {
		ids = append(ids, k)
		vectors = append(vectors, v)
	}
	df := dataframe.New(series.New(vectors[0], series.Float, strconv.Itoa(int(ids[0]))))

	for i, v := range vectors {
		if i == 0 {
			continue
		}
		df2add := dataframe.New(series.New(v, series.Float, strconv.Itoa(int(ids[i]))))
		df = df.CBind(df2add)
	}

	df = df.Filter(
		dataframe.F{Colidx: 0, Comparator: series.Greater, Comparando: 1e-4},
	)
	vectors32 := make([][]float32, 0)

	for _, id := range ids {
		vector := df.Col(strconv.Itoa(int(id))).Float()
		targetVector := make([]float32, len(vector))
		varItem := vector[0]
		varValue := 0.
		for i, v := range vector {
			targetVector[i] = float32(v)
			if math.Abs(v-varItem) > varValue {
				varValue = math.Abs(v - varItem)
				varItem = v
			}
		}
		vectors32 = append(vectors32, targetVector)
	}
	dim := df.Nrow()
	m.prepare(ctx, int64(dim))

	idColumn := entity.NewColumnInt64("ID", ids)
	vectorColumn := entity.NewColumnFloatVector("Vector", dim, vectors32) //entity.NewColumnFloatVector("Vector", dim, vectors)
	if _, err := m.client.Insert(ctx, collectionName, "", idColumn, vectorColumn); err != nil {
		return fmt.Errorf("failed to insert tick1day data, %w", err)
	}

	if err := m.client.Flush(ctx, collectionName, false); err != nil {
		return fmt.Errorf("failed to flush collection: %w", err)
	}

	idx, err := entity.NewIndexIvfFlat(entity.L2, 2)
	if err != nil {
		log.Fatal("fail to create ivf flat index:", err.Error())
	}
	err = m.client.CreateIndex(ctx, collectionName, "Vector", idx, false)
	if err != nil {
		log.Fatal("fail to create index:", err.Error())
	}

	return nil
}

func (m *Milvus) Close() error {
	if m.client == nil {
		return nil
	}
	// clean up
	//_ = c.DropCollection(ctx, collectionName)
	if err := m.client.Close(); err != nil {
		return fmt.Errorf("failed to close client: %w", zap.Error(err))
	}
	m.client = nil

	return nil
}

func dateToId(timestamp int64, startDate time.Time) int64 {
	date := time.UnixMilli(timestamp)
	diff := date.Sub(startDate)
	days := int64(diff.Hours() / 24)

	return days
	//return int64(date.Year())*10000 + int64(date.Month())*100 + int64(date.Day())
}

/*
func main() {

	// setup context for client creation, use 2 seconds here
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()

	c, err := client.NewClient(ctx, client.Config{
		Address: milvusAddr,
	})

	if err != nil {
		// handling error and exit, to make example simple here
		log.Fatal("failed to connect to milvus:", err.Error())
	}
	// in a main func, remember to close the client
	defer c.Close()

	// here is the collection name we use in this example
	//collectionName := `gosdk_insert_example`

	has, err := c.HasCollection(ctx, collectionName)
	if err != nil {
		log.Fatal("failed to check whether collection exists:", err.Error())
	}
	if has {
		// collection with same name exist, clean up mess
		_ = c.DropCollection(ctx, collectionName)
	}

	// define collection schema, see tick1day.csv
	schema := entity.NewSchema().WithName(collectionName).WithDescription("`Stock exchange daily ticks`").
		WithField(entity.NewField().WithName("ID").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true)).
		WithField(entity.NewField().WithName("StockId").WithDataType(entity.FieldTypeInt64)).
		WithField(entity.NewField().WithName("Vector").WithDataType(entity.FieldTypeFloatVector).WithDim(8))

	err = c.CreateCollection(ctx, schema, entity.DefaultShardNumber) // only 1 shard
	if err != nil {
		log.Fatal("failed to create collection:", err.Error())
	}

	films, err := loadTicksPerDay()
	if err != nil {
		log.Fatal("failed to load tick1day data csv:", err.Error())
	}

	// row-base covert to column-base
	ids := make([]int64, 0, len(films))
	years := make([]int32, 0, len(films))
	vectors := make([][]float32, 0, len(films))
	// string field is not supported yet
	idTitle := make(map[int64]string)
	for idx, film := range films {
		ids = append(ids, film.ID)
		idTitle[film.ID] = film.Title
		years = append(years, film.Year)
		vectors = append(vectors, films[idx].Vector[:]) // prevent same vector
	}
	idColumn := entity.NewColumnInt64("ID", ids)
	yearColumn := entity.NewColumnInt32("Year", years)
	vectorColumn := entity.NewColumnFloatVector("Vector", 8, vectors)

	// insert into default partition
	_, err = c.Insert(ctx, collectionName, "", idColumn, yearColumn, vectorColumn)
	if err != nil {
		log.Fatal("failed to insert tick1day data:", err.Error())
	}
	log.Println("insert completed")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	err = c.Flush(ctx, collectionName, false)
	if err != nil {
		log.Fatal("failed to flush collection:", err.Error())
	}
	log.Println("flush completed")

	// load collection with async=false
	err = c.LoadCollection(ctx, collectionName, false)
	if err != nil {
		log.Fatal("failed to load collection:", err.Error())
	}
	log.Println("load collection completed")

	searchFilm := films[0] // use first fim to search
	vector := entity.FloatVector(searchFilm.Vector[:])
	// Use flat search param
	sp, _ := entity.NewIndexFlatSearchParam()
	sr, err := c.Search(ctx, collectionName, []string{}, "Year > 1990", []string{"ID"}, []entity.Vector{vector}, "Vector",
		entity.L2, 10, sp)
	if err != nil {
		log.Fatal("fail to search collection:", err.Error())
	}
	for _, result := range sr {
		var idColumn *entity.ColumnInt64
		for _, field := range result.Fields {
			if field.Name() == "ID" {
				c, ok := field.(*entity.ColumnInt64)
				if ok {
					idColumn = c
				}
			}
		}
		if idColumn == nil {
			log.Fatal("result field not math")
		}
		for i := 0; i < result.ResultCount; i++ {
			id, err := idColumn.ValueByIdx(i)
			if err != nil {
				log.Fatal(err.Error())
			}
			title := idTitle[id]
			fmt.Printf("file id: %d title: %s scores: %f\n", id, title, result.Scores[i])
		}
	}

	// clean up
	_ = c.DropCollection(ctx, collectionName)

}*/
