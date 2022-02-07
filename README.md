# BarRadio
BARRADIO - приложение для выбора жанра музыки, играющей в общественном заведении по предпочтениям большинства посетиттелей.   

## How it works
При заходе в бар посетитель подключается к сети Wi-Fi, где выполняет своего рода check-in, тем самым разрешая API приложения доступ к предпочитаемому жанру (т.е. либо подключение к установленному стриминговому сервису, в котором определен жанр, либо при чекине выбирает жанр/настроение, который ему хотелось бы сейчас слушать). Пользователь добавляется в локальную БД (либо в память самого приложнеия) с названием его желаемого жанра. Производится подсчет совпадений во вкусах по таблице и на основе этого составляется плей-лист для воспроизведения

## Начало работы ##

Для ознакомления можно воспользоваться [демо-сервером](https://barradio.herokuapp.com/api/hello).

Спецификацию на API вы можете найти [тут]().

## Установка ##

TODO...

## Версионирование ##

Мы используем [SemVer](http://semver.org/). Доступные версии можно увидеть  [в тегах репозитория](https://github.com/larikhide/barradio/tags). 

## TODO
- определить минимальный состав проекта. Что в нем должно быть? (сервер, клиент, БД, api)
- как можно определить предпочитаемый жанр? (может как-то под-другому от того, что описано в описании? )
- какие разделы добавить в README? (технологии, которые используются, как запустить приложение)

## Cписок участников
P. Ustyuzhanin https://github.com/larikhide  
А. Куприянов https://github.com/sanya-spb  
Дмитрий https://github.com/shda1615  
Vitaliy Shatskikh https://github.com/vitalyshatskikh  
А. Афанасьев https://github.com/lewa100  
D. Bulavin https://github.com/DENDarkness  
D. Alekseev https://github.com/alekceev
Vladimir https://github.com/fladago
