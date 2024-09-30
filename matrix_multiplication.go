go
package main

import (
    "fmt"
    "sync"
)

// Функция для умножения частичной матрицы
func multiplyRow(matrixA [][]int, matrixB [][]int, result [][]int, row int, wg *sync.WaitGroup) {
    defer wg.Done() // Уменьшение счетчика ожидания по завершении функции
    for j := 0; j < len(matrixB[0]); j++ {
        sum := 0
        for k := 0; k < len(matrixB); k++ {
            sum += matrixA[row][k] * matrixB[k][j]
        }
        result[row][j] = sum
    }
}

// Функция для перемножения матриц
func multiplyMatrices(matrixA [][]int, matrixB [][]int) [][]int {
    // Получаем размерности
    rowsA := len(matrixA)
    colsA := len(matrixA[0])
    colsB := len(matrixB[0])

    // Проверка совместимости матриц
    if colsA != len(matrixB) {
        panic("Матрицы имеют несовместимые размеры")
    }

    // Инициализация результирующей матрицы
    result := make([][]int, rowsA)
    for i := range result {
        result[i] = make([]int, colsB)
    }

    var wg sync.WaitGroup

    // Запускаем горутины для умножения каждого ряда
    for i := 0; i < rowsA; i++ {
        wg.Add(1) // Увеличиваем счетчик ожидания
        go multiplyRow(matrixA, matrixB, result, i, &wg)
    }

    wg.Wait() // Ожидание завершения всех горутин

    return result
}

// Функция для отображения матрицы
func printMatrix(matrix [][]int) {
    for _, row := range matrix {
        fmt.Println(row)
    }
}

func main() {
    // Пример матриц для перемножения
    matrixA := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    matrixB := [][]int{
        {7, 8},
        {9, 10},
        {11, 12},
    }

    // Перемножение матриц
    result := multiplyMatrices(matrixA, matrixB)

    // Вывод результата
    fmt.Println("Результирующая матрица:")
    printMatrix(result)
}
