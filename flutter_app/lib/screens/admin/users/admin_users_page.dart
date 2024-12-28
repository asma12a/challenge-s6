import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/services/user_service.dart';
import 'add_edit_user_page.dart';
import 'dart:ui';

class AdminUsersPage extends StatefulWidget {
  const AdminUsersPage({super.key});

  @override
  State<AdminUsersPage> createState() => _AdminUsersPageState();
}

class _AdminUsersPageState extends State<AdminUsersPage> {
  late Future<List<UserApp>> _usersFuture;

  @override
  void initState() {
    super.initState();
    _loadUsers();
  }

  void _loadUsers() {
    setState(() {
      _usersFuture = UserService.getUsers();
    });
  }

  @override
  Widget build(BuildContext context) {
    final authState = context.read<AuthState>();

    if (!authState.isAdmin) {
      return Scaffold(
        appBar: AppBar(
          title: const Text('Accès interdit'),
        ),
        body: const Center(
          child: Text(
              'Vous n\'avez pas les permissions nécessaires pour accéder à cette page.'),
        ),
      );
    }

    return Scaffold(
      appBar: AppBar(
        title: const Text('Gestion des utilisateurs'),
      ),
      body: FutureBuilder<List<UserApp>>(
        future: _usersFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(
              child: Text(
                'Erreur: ${snapshot.error}',
                style: TextStyle(color: Theme.of(context).colorScheme.error),
              ),
            );
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(child: Text('Aucun utilisateur trouvé.'));
          }

          final users = snapshot.data!;

          return UsersList(
            users: users,
            onRefresh: _loadUsers, // Passer la fonction de rafraîchissement
          );
        },
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Affichage du modal avec flou d'arrière-plan
          _showAddEditUserDialog(context);
        },
        child: const Icon(Icons.add),
      ),
    );
  }

  // Fonction pour afficher le formulaire dans un dialog avec arrière-plan flouté
  void _showAddEditUserDialog(BuildContext context) {
    showDialog(
      context: context,
      barrierDismissible:
          true, // Permet de fermer le modal en cliquant en dehors
      builder: (BuildContext context) {
        return Dialog(
          backgroundColor:
              Colors.transparent, // Transparent pour laisser passer le flou
          child: BackdropFilter(
            filter: ImageFilter.blur(sigmaX: 5.0, sigmaY: 5.0),
            child: Material(
              color:
                  Colors.transparent, // Transparent pour laisser passer le flou
              child: Center(
                // Centre le contenu de la modal
                child: Container(
                  padding: const EdgeInsets.all(16.0),
                  constraints: BoxConstraints(
                    maxWidth: 600, // Limite la largeur maximale de la modal
                  ),
                  decoration: BoxDecoration(
                    color: Colors.white, // Fond blanc pour la modal
                    borderRadius: BorderRadius.circular(12), // Bordure arrondie
                    border: Border.all(
                      color: Colors
                          .transparent, // Bordure transparente (ou modifiez la couleur)
                      width: 1,
                    ),
                  ),
                  child: Stack(
                    children: [
                      Padding(
                        padding: const EdgeInsets.all(16.0),
                        child: AddEditUserPage(), // Le formulaire ici
                      ),
                      Positioned(
                        top: 16,
                        right: 16,
                        child: IconButton(
                          icon: const Icon(Icons.close,
                              color: Colors.black, size: 30), // Croix noire
                          onPressed: () {
                            Navigator.of(context).pop(); // Ferme la modal
                          },
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ),
          ),
        );
      },
    );
  }
}

class UsersList extends StatelessWidget {
  final List<UserApp> users;
  final VoidCallback onRefresh;

  const UsersList({Key? key, required this.users, required this.onRefresh})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: users.length,
      itemBuilder: (context, index) {
        final user = users[index];

        return Card(
          key: ValueKey(user.id),
          margin: const EdgeInsets.symmetric(vertical: 8, horizontal: 16),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(8),
          ),
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  user.name,
                  style: Theme.of(context).textTheme.titleMedium?.copyWith(
                        fontWeight: FontWeight.bold,
                      ),
                ),
                const SizedBox(height: 8),
                Text(user.email, style: Theme.of(context).textTheme.bodyLarge),
                const SizedBox(height: 4),
                Text(
                  user.roles.map((role) => role.name).join(', '),
                  style: Theme.of(context).textTheme.bodyMedium?.copyWith(
                        color: Theme.of(context).colorScheme.primary,
                      ),
                ),
                const Divider(height: 20),
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: [
                    IconButton(
                      icon: const Icon(Icons.edit, color: Colors.blue),
                      onPressed: () async {
                        await Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) => AddEditUserPage(user: user),
                          ),
                        );
                        onRefresh(); // Rafraîchir après modification
                      },
                    ),
                    IconButton(
                      icon: const Icon(Icons.delete, color: Colors.red),
                      onPressed: () {
                        _confirmDelete(context, user);
                      },
                    ),
                  ],
                ),
              ],
            ),
          ),
        );
      },
    );
  }

  void _confirmDelete(BuildContext context, UserApp user) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Supprimer utilisateur'),
          content: Text('Voulez-vous vraiment supprimer ${user.email} ?'),
          actions: [
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              child: const Text('Annuler'),
            ),
            TextButton(
              onPressed: () async {
                await UserService.deleteUser(user.id); // Suppression réelle
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('${user.email} a été supprimé.')),
                );
                Navigator.of(context).pop();
                onRefresh(); // Rafraîchir la liste après suppression
              },
              child: const Text('Confirmer'),
            ),
          ],
        );
      },
    );
  }
}
