# VtbApiHack
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)![Go](https://img.shields.io/badge/golang-%23007ACC.svg?style=for-the-badge&logo=go&logoColor=white)![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)![Figma](https://img.shields.io/badge/figma-%2320232a.svg?style=for-the-badge&logo=figma)



# [Ссылка на готовое решение](http://vtb.shmyaks.ru/)
# [Figma](https://www.figma.com/design/NKToYQvFZ7sKCCjoI6aTYq/vtb-api?node-id=0-1&t=oSV0RNYPFzDoB8hU-1)

### Задача: Создание приложений на основе Open Banking API

## Используемый стек технологий:
- [GO-Backend](https://github.com/ultraevs/VtbApiHack/tree/main/go-backend) - Реализован с использванием [GO](https://go.dev/) и фреймворка [Gin](https://github.com/gin-gonic/gin). Задачей модуля является реализация API для взаимодействия с frontend модулем.
- [Frontend](https://github.com/ultraevs/VtbApiHack/tree/main/frontend) - Реализован с использованием [React](https://ru.legacy.reactjs.org/). Задачай является предоставление красивого и функционалоного интерфейса для пользователя.
- [Deployment](https://github.com/ultraevs/VtbApiHack/tree/main/deployment) - Реализован с использованием [Docker-Compose](https://www.docker.com/). Задачей модуля является возможность быстрого и безошибочного развертывания приложения на любом сервере.

## Функционал решения

- Личный аккаунт.
- История переводов.
- Раздел карт, кредитов, инвестиций.
- Взаимодействие с open banking api.

## Запуск решения
```sh
    cd VtbApiHack/deployment
    docker-compose build
    docker-compose up -d
```

##  Состав команды
 - Михаил Евсеев - Backend Developer
 - Артем Брежнев - Designer
 - Иван Лобода - Backend Developer
 - Костя Цой - Frontend Developer
 - Глеб Таланцев - Frontend Developer и Project Manager

## Итог
Проект VtbApiHack представляет собой современное и интуитивно понятное приложение для управления банковскими услугами. Основная цель решения — предоставить пользователям простой доступ к информации о своих финансах, включая карты, кредиты и инвестиции, а также реализовать функционал для интеграции с внешними Open Banking API.

Проект выделяется своим удобным интерфейсом, быстрым и безопасным бекендом, а также возможностью масштабируемости за счет использования современных технологий. Интеграция с открытыми банковскими API делает его универсальным инструментом, подходящим для пользователей с различными финансовыми потребностями.

Отличия и важность

Удобство и безопасность: Возможность управления финансами в одном месте с защитой данных на всех этапах взаимодействия.
Интеграции: Работа с Open Banking API позволяет подключать сторонние сервисы для более широкого функционала.
Масштабируемость: Архитектура на основе Docker и микросервисов позволяет легко адаптировать приложение под новые задачи.

## Возможные улучшения
Расширение функционала:

Поддержка платежей через QR-коды и NFC.
Добавление персонализированных рекомендаций для пользователей на основе их транзакций.
Технические доработки:

Внедрение аналитики и дашбордов для отслеживания финансовых показателей.
Интеграция с дополнительными банковскими и инвестиционными сервисами.
Улучшение UX/UI:

Поддержка мобильного приложения для удобного управления финансами в пути.
Внедрение темной темы для улучшения пользовательского опыта.
Автоматизация процессов:

Уведомления о сроках погашения кредитов, завершении депозитов или значительных движениях по счету.
Возможность автоматического формирования отчетов по финансам.
