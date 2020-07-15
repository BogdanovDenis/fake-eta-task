
## Command line flags

Настройки:

* *--car_url* - url сервиса Car
* *--predict_url* - url сервиса Predict

## Build
Для кодогенерации был использован go-swagger v0.19.0

```sh
cd fake-eta-task/
swagger generate server -A wait -P models.Principal -f swagger/server/wait.yml   
swagger generate client -A car -c internal/car -m internal/car/models -f swagger/clients/car.yml  
swagger generate client -A predict -c internal/predict -m internal/predict/models -f swagger/clients/predict.yml  
cd cmd/wait-server/
go build
./wait-server --car-url=127.0.0.1:53726 --predict-url=127.0.0.1:53737
```