package main

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert into to ByteCounter，用原生int的+操作
	return len(p), nil
}

func (c *ByteCounter) Get() int {
	return int(*c)
}
