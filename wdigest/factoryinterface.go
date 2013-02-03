package main

type FactoryInterface interface {
    TokenFactory
    TokenContainerFactory
    MessagerFactory
    MessagerContainerFactory
}
