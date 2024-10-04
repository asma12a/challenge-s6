# Challenge 5IWJ-S2

#### Frontend: Flutter

#### Backend: Golang

- Framework Web: [Fiber](https://gofiber.io/)
- ORM: [Ent](https://entgo.io/)

## Commandes utiles

### Golang

```
docker compose up -d
```

Lancer le serveur (watcher [air](https://github.com/air-verse/air))

```
air
```

Migration de la base de données

```
go run database/migrate/main.go
```

Créer un modèle

```
go run entgo.io/ent/cmd/ent new <nom-de-modele>
```

Après avoir édité le modèle:

```
go generate ./ent
```

### Flutter

Lancer l'application

```
flutter run
```
