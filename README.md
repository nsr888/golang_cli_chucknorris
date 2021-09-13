![Go Report](https://goreportcard.com/badge/github.com/nsr888/golang_cli_chucknorris)
![Repository Top Language](https://img.shields.io/github/languages/top/nsr888/golang_cli_chucknorris)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/nsr888/golang_cli_chucknorris)
![Github Repository Size](https://img.shields.io/github/repo-size/nsr888/golang_cli_chucknorris)
![Lines of code](https://img.shields.io/tokei/lines/github/nsr888/golang_cli_chucknorris)
![GitHub last commit](https://img.shields.io/github/last-commit/nsr888/golang_cli_chucknorris)

# Golang CLI-приложение для работы с https://api.chucknorris.io/

Необходимо реализовать CLI-приложение для работы с https://api.chucknorris.io/. Приложение должно уметь: печатать на экран случайную шутку; выгружать по n случайных уникальных шуток ( n задается пользователем, дефолтное значение 5) для каждой из существующих категорий и сохранять их в текстовые файлы - по одному на каждую из категорий.

Пример Получение случайной шутки: Выгрузка шуток по категориям:

Требования решение должно быть в git-репозитории (можно прислать архив или опубликовать на github, gitlab, bitbucket...). Пожелания документирование кода; тесты; использование статического анализатора (конфигурацию положить в репозиторий).

```
$ joker random
Chuck Norris`s kicks are so fast ..you probably wudnt feel it till yesterday.
$ joker dump -n 2
$ ls -l
animal.txt career.txt celebrity.txt ...
$ cat animal.txt
Chuck Norris once rode a bull, and nine months later it had a calf. Chuck Norris does not own a house. He walks into random houses and people move
```
