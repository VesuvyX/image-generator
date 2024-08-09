# GENERATOR IMAGE SERVER

## STACK

Golang

## FUNCTIONAL

### RU

С помощью адресной строки вы можете создавать блок определенного размера, цвета и с текстом внутри (Тексту можно задавать размер)
http://localhost:8080/Ширина/Высота/ЦветБлока/Текст/ЦветТекста/РазмерТекста (Если не указан текст, то блок создается без текста)

### ENG

Using the address bar you can create a block of a specific size, color and text inside (Text can be set in size)
http://localhost:8080/Width/Height/BlockColor/Text/TextColor/TextSize (If no text is specified, the block is created without text)

## Instructions for running code locally

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/VesuvyX/image-generator/
   ```
2. Перейдите в директорию проекта:

   ```bash
   cd image-generator
   ```

3. Запустите сервер:
   ```bash
   go run cmd/main.go
   ```
4. Откройте браузер и введите адрес:
   ```
   http://localhost:8080
   ```
