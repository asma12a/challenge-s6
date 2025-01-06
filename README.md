# Challenge 5IWJ-S2

#### Backend: Golang

- Framework Web: [Fiber](https://gofiber.io/)
- ORM: [Ent](https://entgo.io/)

## 🚀 Procédure d'installation et de lancement

1- Créer un fichier `.env` au même niveau que le `.env.example` et dupliquer son contenu dedans.

2 - Lancer le `docker-compose` en mode détaché et s'assurer que les 3 containers sont up :
    ```bash
    docker compose up -d
    ```

3 - Lancer le serveur (avec le watcher [air](https://github.com/air-verse/air)) :
    ```bash
    air
    ```

4 - Migration de la base de données :
    ```bash
    go run database/migrate/main.go
    ```

5 - Créer un modèle :
    ```bash
    go run entgo.io/ent/cmd/ent new <nom-de-modele>
    ```

6 - Après avoir édité le modèle, exécuter :
    ```bash
    go generate ./ent
    ```

### 📜 Documentation Swagger

Une fois que l'application backend est en cours d'exécution, tu peux accéder à la documentation interactive de l'API via **Swagger UI**. Cela te permettra de découvrir et tester facilement les différentes routes de l'API.

1. Démarre ton serveur backend avec la commande suivante :

    ```bash
    air
    ```

2. Une fois l'application démarrée, lance la commande **go run server.go**, ouvre ton navigateur et rends-toi à l'URL suivante pour accéder à **Swagger UI** :

    ```
    http://localhost:3001/swagger/index.html
    ```

3. Dans Swagger UI, tu pourras explorer toutes les routes disponibles, avec la possibilité de tester les différentes requêtes directement depuis l'interface web.

#### Frontend: Flutter

Lancer l'application :
    ```bash
    flutter run
    ```
