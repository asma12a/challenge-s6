import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/services/user_service.dart';
import 'add_edit_user_page.dart';
import '../../custom_data_table.dart';
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

          final columns = [
            DataColumn(label: Text('Nom')),
            DataColumn(label: Text('Email')),
            DataColumn(label: Text('Rôles')),
            DataColumn(label: Text('Actions')),
          ];

          final rows = users.map((user) {
            return DataRow(cells: [
              DataCell(Text(user.name)),
              DataCell(Text(user.email)),
              DataCell(Text(user.roles.map((role) => role.name).join(', '))),
              DataCell(Row(
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
                      _loadUsers(); // Rafraîchir après modification
                    },
                  ),
                  IconButton(
                    icon: const Icon(Icons.delete, color: Colors.red),
                    onPressed: () {
                      _confirmDelete(context, user);
                    },
                  ),
                ],
              )),
            ]);
          }).toList();

          return CustomDataTable(
            title: 'Utilisateurs',
            columns: columns,
            rows: rows,
          );
        },
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          _showAddEditUserDialog(context);
        },
        child: const Icon(Icons.add),
      ),
    );
  }

  void _showAddEditUserDialog(BuildContext context) {
    showDialog(
      context: context,
      barrierDismissible: true,
      builder: (BuildContext context) {
        return Dialog(
          backgroundColor: Colors.transparent,
          child: BackdropFilter(
            filter: ImageFilter.blur(sigmaX: 5.0, sigmaY: 5.0),
            child: Material(
              color: Colors.transparent,
              child: Center(
                child: Container(
                  padding: const EdgeInsets.all(16.0),
                  constraints: BoxConstraints(maxWidth: 600),
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(12),
                    border: Border.all(color: Colors.transparent, width: 1),
                  ),
                  child: Stack(
                    children: [
                      Padding(
                        padding: const EdgeInsets.all(16.0),
                        child: AddEditUserPage(),
                      ),
                      Positioned(
                        top: 16,
                        right: 16,
                        child: IconButton(
                          icon: const Icon(Icons.close,
                              color: Colors.black, size: 30),
                          onPressed: () {
                            Navigator.of(context).pop();
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

  void _confirmDelete(BuildContext context, UserApp user) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Supprimer l\'utilisateur'),
          content: Text('Voulez-vous vraiment supprimer ${user.email} ?'),
          actions: [
            // Bouton Annuler - couleur gris et texte noir
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              child: const Text(
                'Annuler',
                style:
                    TextStyle(color: Colors.black), // Texte noir pour Annuler
              ),
            ),
            // Bouton Confirmer - avec bordure rouge et texte blanc
            TextButton(
              onPressed: () async {
                await UserService.deleteUser(user.id);
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('${user.email} a été supprimé.')),
                );
                Navigator.of(context).pop();
                _loadUsers(); // Rafraîchir après suppression
              },
              style: TextButton.styleFrom(
                side: BorderSide(color: Colors.red, width: 2), // Bordure rouge
                backgroundColor: Colors.red, // Fond rouge
                padding: EdgeInsets.symmetric(
                    vertical: 12, horizontal: 24), // Espacement intérieur
              ),
              child: const Text(
                'Confirmer',
                style: TextStyle(
                    color: Colors.white), // Texte blanc pour Confirmer
              ),
            ),
          ],
        );
      },
    );
  }
}
