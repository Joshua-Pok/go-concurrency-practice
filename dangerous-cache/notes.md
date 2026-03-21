<!--markdownlint-disable-->


# Step 1: Sequential Store


## Step Goal: Implement a store struct that satisfies a Cache interface with three methods: Get, Set, Delete


It should store data in a standard go map


Sketch

type Cache[T any] Interface{


Get() T
Set() void
Delete() void

}


type store[int cache] struct

# Dependency Inversion

By defining Cache interface first, we ensure the rest of the application dosent care how the cache works, only that it satisfies the contract


Normally a high level piece of code (function that caculates a user total rewards) would directly create and use a low level tool (database client)


this is bad because if we change the db we need to rewrite the rewards logic


DIP flips this by saying that **high level modules should not depend on low level modules.**

**Abstractions should not depend on details, details should depend on abstractions**



sketch 

type Cache[T any] interface{

Get(key string)(T, error)
Set(key string ,value T) error
Delete(key string)error
}

type Store[T any] struct{
data map[string]T
}

func NewStore[T any]() Cache[T]{
return &Store[T]{
data: make(map[string]T),
}
}
