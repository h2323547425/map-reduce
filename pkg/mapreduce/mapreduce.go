package mapreduce

// Mapper represents a map function.
type Mapper[I any, K comparable, V any] func(input I) map[K]V

// Reducer represents a reduce function.
type Reducer[K comparable, V, O any] func(key K, values []V) O

// Job represents a MapReduce job.
type Job[I any, K comparable, V, O any] struct {
	inputs  []I
	mapper  Mapper[I, K, V]
	reducer Reducer[K, V, O]
}

// Map function applies the mapper to each input value.
func (j *Job[I, K, V, O]) Map() []map[K]V {
	mapResults := make([]map[K]V, 0)
	emitChan := make(chan map[K]V)
	for _, input := range j.inputs {
		go func(input I) {
			emitChan <- j.mapper(input)
		}(input)
	}
	for i := 0; i < len(j.inputs); i++ {
		mapResults = append(mapResults, <-emitChan)
	}
	return mapResults
}

// Reduce function applies the reducer to the output of the mapper.
func (j *Job[I, K, V, O]) Reduce(mapResult map[K][]V) map[K]O {
	reduceResult := make(map[K]O)
	type Result struct {
		Key    K
		Result O
	}
	emitChan := make(chan Result)
	for key, values := range mapResult {
		go func(key K, values []V) {
			emitChan <- Result{key, j.reducer(key, values)}
		}(key, values)
	}
	for i := 0; i < len(mapResult); i++ {
		result := <-emitChan
		reduceResult[result.Key] = result.Result
	}
	return reduceResult
}

// Run function runs the MapReduce job.
func (j *Job[I, K, V, O]) Run() map[K]O {
	mapResult := make(map[K][]V)
	for _, result := range j.Map() {
		for key, value := range result {
			if _, ok := mapResult[key]; !ok {
				mapResult[key] = make([]V, 0)
			}
			mapResult[key] = append(mapResult[key], value)
		}
	}
	return j.Reduce(mapResult)
}

func MapReduce[I any, K comparable, V, O any](inputs []I, mapper Mapper[I, K, V], reducer Reducer[K, V, O]) map[K]O {
	job := &Job[I, K, V, O]{
		inputs:  inputs,
		mapper:  mapper,
		reducer: reducer,
	}
	return job.Run()
}
