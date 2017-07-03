package main

func main() {
	RegisterPrometheus()
	gauge.WithLabelValues("label1", "label2").Inc()
}
