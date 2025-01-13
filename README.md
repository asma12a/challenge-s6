# Challenge 5IWJ-S2

## 🚀 Backend : Golang

Le backend est développé en **Golang** et se concentre sur la gestion des événements sportifs et des utilisateurs, avec des fonctionnalités avancées comme le WebSocket pour le chat en temps réel.

### 📌 Technologies principales utilisées

- **Framework Web** : [Fiber](https://gofiber.io/)  
- **ORM** : [Ent](https://entgo.io/) pour interagir avec PostgreSQL  
- **Base de données** : PostgreSQL  
- **Cache** : Redis pour gérer les données éphémères, comme les codes de confirmation  
- **WebSocket** : pour la communication en temps réel (chat des événements)  
- **Documentation API** : Swagger UI  

---

### 📋 Fonctionnalités principales

1. **API REST** :
   - CRUD complet pour les entités principales : événements, utilisateurs, équipes, rôles, etc.
   - Routes sécurisées avec gestion des permissions basée sur les rôles.  

2. **Système d'authentification** :
   - Inscription avec email et mot de passe.
   - Envoi d'un email de confirmation contenant un code valable 15 minutes (stocké dans Redis).
   - Gestion des tokens JWT pour sécuriser les appels API.

3. **Gestion des rôles** :
   - Rôles disponibles : Administrateur, Organisateur, Coach.
   - Contrôle d'accès pour chaque route, avec filtrage des données selon le rôle.  

4. **WebSocket** : **Chat en temps réel**
   - Mise en place d'un système de chat en temps réel permettant aux utilisateurs de discuter au sein des événements.  
   - Fonctionnalités principales :  
     - Joindre une salle de chat correspondant à un événement spécifique.  
     - Envoyer et recevoir des messages en temps réel.  
     - Gestion des connexions multiples, avec détection des déconnexions.  
   - Exemple d'implémentation :  
     - Endpoint WebSocket pour rejoindre le chat : `/ws/chat/:eventID`.  
     - Gestion des salles basées sur `eventID` pour isoler les discussions par événement.  

5. **Swagger UI** :
   - Accessible via l'URL suivante :  

     ```
     http://localhost:3001/swagger/index.html
     ```

   - Permet d'explorer et de tester facilement toutes les routes de l'API.

6. **Pagination** :
   - Pagination pour toutes les routes `GET all`, configurable via des paramètres de requête (`page`, `limit`).

7. **Service tiers** :
   - Intégration de Google Maps pour géolocaliser les événements et afficher leur emplacement.

8. **Logs formatés** :
   - Gestion des logs structurés pour suivre les opérations critiques et déboguer les erreurs.  
   - Utilisation d'une bibliothèque de gestion des logs (par ex. `logrus`).  

9. **Configuration flexible** :
   - Gestion des configurations via un fichier `.env` et variables d'environnement.  
   - Les configurations incluent : informations de la base de données, ports d'écoute, clés JWT, etc.

---

### ⚙️ Procédure d'installation et de lancement

#### Étape 1 : Configurer les variables d'environnement

Créez un fichier `.env` au même niveau que le fichier `.env.example` et personnalisez son contenu selon votre environnement. Exemple :  

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=sports_app
JWT_SECRET=myjwtsecret
REDIS_ADDR=localhost:6379
EMAIL_HOST=smtp.example.com
EMAIL_PORT=587
EMAIL_USER=your_email@example.com
EMAIL_PASSWORD=email_password
Étape 2 : Lancer les services Docker
```

Assurez-vous que Docker est installé. Lancez les containers nécessaires avec :
bash
Copier le code
docker compose up -d
Cela démarre :

PostgreSQL pour la base de données.
Redis pour le cache.
Étape 3 : Lancer le serveur
Utilisez air pour démarrer le serveur en mode développement :

bash
Copier le code
air
Étape 4 : Migration de la base de données
Appliquez les migrations de la base de données pour initialiser les tables :

bash
Copier le code
go run database/migrate/main.go
Étape 5 : Créer un modèle
Pour ajouter de nouvelles entités au projet, utilisez la commande suivante :

bash
Copier le code
go run entgo.io/ent/cmd/ent new <nom-du-modele>
Après avoir édité le modèle, générez le code avec :

bash
Copier le code
go generate ./ent
📜 Exemples d'utilisation
Swagger UI
Une fois le serveur backend lancé, accédez à la documentation interactive à :

bash
Copier le code
http://localhost:3001/swagger/index.html
WebSocket Chat
Connectez-vous au WebSocket pour un événement spécifique :
ruby
Copier le code
ws://localhost:3001/ws/chat/:eventID
Exemple de message envoyé au serveur (format JSON) :
json
Copier le code
{
  "username": "JohnDoe",
  "message": "Hello, team!"
}
Les autres utilisateurs dans la salle recevront le message en temps réel.
🧪 Tests
Tests unitaires :
Testez les règles de gestion via :
bash
Copier le code
go test ./...
Tests d'intégration :
Vérifiez les routes API et les fonctionnalités principales avec des mocks.
