/*
## Задача № 1
Написать API для указанных маршрутов(endpoints)
"/info"   // Информация об API
"/first"  // Случайное число
"/second" // Случайное число
"/add"    // Сумма двух случайных чисел
"/sub"    // Разность
"/mul"    // Произведение
"/div"    // Деление

*результат вернуть в виде JSON

"math/rand"
number := rand.Intn(100)

// Queries
GET http://127.0.0.1:1234/first

GET http://127.0.0.1:1234/second

GET http://127.0.0.1:1234/add
GET http://127.0.0.1:1234/sub
GET http://127.0.0.1:1234/mul
GET http://127.0.0.1:1234/div
GET http://127.0.0.1:1234/info
*/