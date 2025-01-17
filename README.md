# Challenge 5IWJ-S2

## 🚀 Frontend : Flutter

# SquadGo - Application Mobile de Gestion d'Événements Sportifs

## Description

**SquadGo** est une application mobile qui permet aux utilisateurs de créer, rejoindre et gérer des événements sportifs. Elle offre une plateforme interactive et intuitive pour organiser des événements, communiquer entre participants et suivre les performances sportives.

Développée avec **Flutter** pour le frontend et **Golang** pour le backend, l'application propose une gestion complète des événements sportifs et des rôles des participants, avec des notifications pour garder les utilisateurs informés à chaque étape.

---

## Contributeurs

- **Bastien DIKIADI** (913Bass)
- **Ahamed Mze Taslima** (Taslima-Ahamed-Mze)
- **Asma MOKEDDES** (asma12a)
- **Daniel MANEA** (dan1M)

---

## Fonctionnalités

### 1. **Authentification**

- **Inscription avec confirmation par mail**
- **Connexion**

### 2. **Accueil**

- **Visualisation des événements** auxquels l'utilisateur participe ou a créés : Affichage des événements à venir.
- **Recommandations d'événements** : Les événements recommandés sont basés sur la position géographique de l'utilisateur. Si l'utilisateur refuse de partager sa localisation, une latitude et longitude par défaut (FRANCE) sont attribuées.

### 3. **Recherche**

- **Recherche d'événements** : Recherche filtrée par type (Match ou Training), et limitée aux événements publics. Il est également possible de rechercher par sport, nom ou adresse.
- **Recherche d'événements privés** : Permet la recherche d'un événement privé à l'aide d'un code d'événement.

### 4. **Profil utilisateur**

- **Modification des informations personnelles** : Les utilisateurs peuvent mettre à jour leurs données personnelles.
- **Visualisation des événements liés à l'utilisateur** : Permet de voir tous les événements auxquels l'utilisateur participe ou qu'il a créés.
- **Suivi des performances** : Visualisation des performances de l'utilisateur par sport (ex. : nombre de buts marqués, etc.).

### 5. **Gestion des événements**

- **Création d'événements sportifs** : L'utilisateur peut créer un événement sportif parmi quatre types de sports par défaut : Football, Basketball, Tennis, Running.
- **Rejoindre un événement** : Possibilité de rejoindre un événement existant via un code d'événement ou directement à partir de la page d'événements.
- **Rôles dans l'événement** :
  - **Joueur classique** : Participant standard.
  - **Coach** : Peut noter les participants (uniquement le jour de l'événement) sur la base des statistiques (ex. : buts marqués en football). Le coach peut également créer des équipes et ajouter des joueurs.
  - **Organisateur** : Peut modifier les informations de l'événement, créer des équipes et gérer les joueurs au sein des équipes (ajouter via e-mail, changer de rôle, ou supprimer).
- **Invitation des joueurs** : Les organisateurs peuvent inviter des joueurs à rejoindre un événement, même s'ils ne sont pas inscrits à l'application. Après l'inscription, leurs vrais prénoms apparaissent dans l'équipe.
- **Partage du code d'événement** : Permet de partager un code pour inviter des amis à rejoindre un événement.
- **Chat par événement** : Permet de discuter avec tous les participants d'un événement.

### 6. **Notifications**

- **Notifications push** :
  - Informer un joueur lorsqu'il a été noté par un coach.
  - Rappel à J-1 de l'événement pour les participants.
  - Notification à l'organisateur lorsqu'un invité non inscrit rejoint l'événement après son inscription.

### 7. **Hors-ligne**

- **Accès aux événements consultés hors-ligne** : L'utilisateur peut consulter les événements qu'il a déjà visualisés même sans connexion Internet.
- **Prévention des actions (CRUD) en mode hors-ligne** : Des pop-ups empêchent l'utilisateur d'effectuer des actions lorsqu'il est hors-ligne.

### 8. **Back-Office**

- **Accès réservé à l'admin** : L'admin peut effectuer des actions CRUD (Créer, Lire, Mettre à jour, Supprimer) sur les ressources suivantes :
  - Sports
  - Utilisateurs
  - Critères de notation par sport
- **Visualisation des logs serveur** : L'admin peut consulter les logs de chaque requête effectuée sur le serveur pour une gestion optimale.

### Compte Admin

Pour accéder au back-office, vous pouvez utiliser le compte suivant :

- **Email** : admin@gmail.com
- **Mot de passe** : adminsquad

### Lien du Back-Office

Vous pouvez accéder au back-office à l'adresse suivante : [https://challenge-s6.vercel.app/](https://challenge-s6.vercel.app/)

---

## Technologies utilisées

- **Frontend** : Flutter
- **API Public** : Firebase, API Gouv, WebSockets, Intl (Traduction)

---

## Installation

1. Clonez ce repository :

   ```bash
   git clone https://github.com/asma12a/challenge-s6.git
   ```

2. Installez les dépendances Flutter :

   ```bash
   flutter pub get
   ```

3. Exécutez l'émulateur avec la commande suivante :
   ```bash
   flutter run --dart-define=API_BASE_URL=https://challenge-s6-1.onrender.com/ --dart-define=JWT_STORAGE_KEY=squadgo-jwt
   ```

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

     ```
      docker compose up -d
     ```

Cela démarre :

PostgreSQL pour la base de données.
Redis pour le cache.
Étape 3 : Lancer le serveur
Utilisez air pour démarrer le serveur en mode développement :

     ```
     air
     ```

Étape 4 : Migration de la base de données

Appliquez les migrations de la base de données pour initialiser les tables :

     ```
      go run database/migrate/main.go
     ```

Étape 5 : Créer un modèle

Pour ajouter de nouvelles entités au projet, utilisez la commande suivante :

     ```
      go run entgo.io/ent/cmd/ent new <nom-du-modele>
     ```

Après avoir édité le modèle, générez le code avec :

     ```
      go generate ./ent
     ```

📜 Exemples d'utilisation
Swagger UI
Une fois le serveur backend lancé, accédez à la documentation interactive à :

     ```
      http://localhost:3001/swagger/index.html>
     ```

WebSocket Chat
Connectez-vous au WebSocket pour un événement spécifique :
` ws://localhost:3001/ws/chat/:eventID`

Exemple de message envoyé au serveur (format JSON) :

```
{
  "username": "JohnDoe",
  "message": "Hello, team!"
}
```

Les autres utilisateurs dans la salle recevront le message en temps réel.

🧪 Tests
Tests unitaires :
Testez les règles de gestion via :
` go test ./...`

Tests d'intégration :
Vérifiez les routes API et les fonctionnalités principales avec des mocks.
