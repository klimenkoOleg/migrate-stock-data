package milvius

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

type Milvius struct {
	config *app.Config
	log    *zap.Logger
	//dim      int
	client client.Client
	data   map[int64][]float64
	dates  map[int64][]string
}

func NewMilvius(config *app.Config, log *zap.Logger) *Milvius {
	return &Milvius{
		config: config,
		log:    log,
		data:   make(map[int64][]float64),
		dates:  make(map[int64][]string),
		//dim:    dim,
	}
}

func (m *Milvius) prepare(ctx context.Context, dim int64) {
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

func (m *Milvius) WriteTick(_ context.Context, ticks []dto.Tick1Day) error {

	// row-base covert to column-base

	//vector := make([]float32, 80) // TODO: use config for 80

	// string field is not supported yet
	//idTitle := make(map[int64]string)
	//data := make(map[int64]float32)

	for _, tick := range ticks {

		vector, ok := m.data[tick.InstrumentId]
		//dates := m.dates[tick.InstrumentId]
		if !ok {
			diff := m.config.DateEnded.Sub(*m.config.DateStarted)
			days := int64(diff.Hours() / 24)
			vector = make([]float64, days)
			//dates = make([]string, days)
		}
		//id := dateToId(tick.Date, *m.config.DateStarted)
		//if math.Abs(tick.Avg) < 1e-4 || m.skipDays[id] {
		//	m.skipDays[id] = true
		//	continue
		//}

		id := dateToId(tick.Date, *m.config.DateStarted)
		vector[id] = float64(tick.Avg / 100_000_000)
		//dates[id] = time.UnixMilli(tick.Date).Format("2006-01-02")
		m.data[tick.InstrumentId] = vector
		//m.dates[tick.InstrumentId] = dates

		//idx := calcIdxStartingFromStartDate(tick.Date, m.config.DateStarted)
		// caring about duplicates in dates
		//data[dateToId(tick.Date, *m.config.DateStarted)] = float32(tick.Avg)
		//vector[dateToId(tick.Date, *m.config.DateStarted)] = float32(tick.Avg)

		//ids = append(ids, dateToId(film.Date))  // use symbol
		//idTitle[film.ID] = film.Title
		//years = append(years, dateToId(film.Date))
		//vectors = append(vectors, float32(film.Avg)) // films[idx].Vector[:]) // prevent same vector
	}
	//m.ids = append(m.ids, ticks[0].InstrumentId)
	//m.vectors = append(m.vectors, vector)

	//for k, v := range data {
	//    m.vectors = append(m.vectors, v)
	//}

	return nil
}

func (m *Milvius) Flush(ctx context.Context) error {
	ids := make([]int64, 0, len(m.data))
	vectors := make([][]float64, 0, len(m.data))
	//dates := make([]string, 0, len(m.dates))

	for k, v := range m.data {
		ids = append(ids, k)
		vectors = append(vectors, v)
		//dates = append(dates, m.dates[k])
	}

	//for _, v := range m.dates /{
	//	dates = append(dates, v)
	//}

	//vectorSeries := make([]series.Series, 0)
	//for i, v := range vectors {
	//	fmt.Println(v)
	//	vectorSeries = append(vectorSeries, series.New(v, series.Float, strconv.Itoa(int(ids[i]))))
	//}

	//dates := series.New(dates, series.String, "dates")
	df := dataframe.New(series.New(vectors[0], series.Float, strconv.Itoa(int(ids[0]))))

	for i, v := range vectors {
		if i == 0 {
			continue
		}
		df2add := dataframe.New(series.New(v, series.Float, strconv.Itoa(int(ids[i]))))
		//fmt.Println(df)
		//fmt.Println(df2add)
		df = df.CBind(df2add)
		//if i == 3 {
		//	break
		//}
	}

	//fmt.Println(df)

	df = df.Filter(
		dataframe.F{Colidx: 0, Comparator: series.Greater, Comparando: 1e-4},
	)

	fmt.Println(df)

	//vectors := map([][]float64, 0)
	//for _, instrumentId := range df.Names() {
	//}

	vectors32 := make([][]float32, 0)

	varMax := 0.
	varMaxId := ids[0]

	for _, id := range ids {
		//ids = append(ids, id)
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

		if varValue > varMax {
			varMax = varValue
			varMaxId = id
		}

		vectors32 = append(vectors32, targetVector)
		//dates = append(dates, m.dates[k])
	}

	fmt.Println("varMaxId!!!", varMaxId)

	fmt.Println(vectors)

	//os.Exit(0)

	//for i := 0; i <
	//
	dim := df.Nrow()
	m.prepare(ctx, int64(dim))

	idColumn := entity.NewColumnInt64("ID", ids)
	//yearColumn := entity.NewColumnInt64("Year", years)
	vectorColumn := entity.NewColumnFloatVector("Vector", dim, vectors32) //entity.NewColumnFloatVector("Vector", dim, vectors)

	//v1 := series.New([]float64{3.0, 4.0}, series.Float, "COL.3")
	//v1.

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

	//sidx := entity.NewScalarIndex()
	//if err := m.client.CreateIndex(ctx, collectionName, "Year", sidx, false); err != nil {
	//	log.Fatal("failed to create scalar index", err.Error())
	//}

	return nil
}

func (m *Milvius) Close() {
	// clean up
	//_ = c.DropCollection(ctx, collectionName)
	if err := m.client.Close(); err != nil {
		m.log.Error("failed to close client: %w", zap.Error(err))
	}
	m.client = nil
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
