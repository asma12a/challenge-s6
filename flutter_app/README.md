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

- **Inscription avec confirmation par mail** : Les utilisateurs peuvent s'inscrire via un formulaire, avec une confirmation de leur adresse e-mail.
- **Connexion** : Les utilisateurs peuvent se connecter à leur compte en utilisant leur adresse e-mail et mot de passe.

### 2. **Page d'accueil**

- **Visualisation des événements** auxquels l'utilisateur participe ou a créés : Affichage des événements à venir.
- **Recommandations d'événements** : Les événements recommandés sont basés sur la position géographique de l'utilisateur. Si l'utilisateur refuse de partager sa localisation, une latitude et longitude par défaut sont attribuées.
- **Recherche d'événements** : Recherche filtrée par type (Match ou Training), et limitée aux événements publics. Il est également possible de rechercher par nom ou adresse.
- **Recherche d'événements privés** : Permet la recherche d'un événement privé à l'aide d'un code d'événement.

### 3. **Profil utilisateur**

- **Modification des informations personnelles** : Les utilisateurs peuvent mettre à jour leurs données personnelles.
- **Visualisation des événements liés à l'utilisateur** : Permet de voir tous les événements auxquels l'utilisateur participe ou qu'il a créés.
- **Suivi des performances** : Visualisation des performances de l'utilisateur par sport (ex. : nombre de buts marqués, etc.).

### 4. **Gestion des événements**

- **Création d'événements sportifs** : L'utilisateur peut créer un événement sportif parmi quatre types de sports par défaut : Football, Basketball, Tennis, Running.
- **Rejoindre un événement** : Possibilité de rejoindre un événement existant via un code d'événement ou directement à partir de la page d'événements.
- **Rôles dans l'événement** :
  - **Joueur classique** : Participant standard.
  - **Coach** : Peut noter les participants (uniquement le jour de l'événement) sur la base des statistiques (ex. : buts marqués en football). Le coach peut également créer des équipes et ajouter des joueurs.
  - **Organisateur** : Peut modifier les informations de l'événement, créer des équipes et gérer les joueurs au sein des équipes (ajouter via e-mail, changer de rôle, ou supprimer).
- **Invitation des joueurs** : Les organisateurs peuvent inviter des joueurs à rejoindre un événement, même s'ils ne sont pas inscrits à l'application. Après l'inscription, leurs vrais prénoms apparaissent dans l'équipe.
- **Partage du code d'événement** : Permet de partager un code pour inviter des amis à rejoindre un événement.
- **Chat par événement** : Permet de discuter avec tous les participants d'un événement.

### 5. **Notifications**

- **Notifications push** :
  - Informer un joueur lorsqu'il a été noté par un coach.
  - Rappel à J-1 de l'événement pour les participants.
  - Notification à l'organisateur lorsqu'un invité non inscrit rejoint l'événement après son inscription.

### 6. **Hors-ligne**

- **Accès aux événements consultés hors-ligne** : L'utilisateur peut consulter les événements qu'il a déjà visualisés même sans connexion Internet.
- **Prévention des actions (CRUD) en mode hors-ligne** : Des pop-ups empêchent l'utilisateur d'effectuer des actions lorsqu'il est hors-ligne.

### 7. **Back-Office**

- **Accès réservé à l'admin** : L'admin peut effectuer des actions CRUD (Créer, Lire, Mettre à jour, Supprimer) sur les ressources suivantes :
  - Sports
  - Utilisateurs
  - Critères de notation par sport
- **Visualisation des logs serveur** : L'admin peut consulter les logs de chaque requête effectuée sur le serveur pour une gestion optimale.

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

3. Configurez le backend Golang (voir la documentation spécifique au backend).

---

## Contribuer

1. Fork ce repository.
2. Créez une branche pour votre fonctionnalité ou correction (`git checkout -b feature/nouvelle-fonctionnalité`).
3. Committez vos modifications (`git commit -am 'Ajout de nouvelle fonctionnalité'`).
4. Poussez votre branche (`git push origin feature/nouvelle-fonctionnalité`).
5. Ouvrez une pull request.

---

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.
