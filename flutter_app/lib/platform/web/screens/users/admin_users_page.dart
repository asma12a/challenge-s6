import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/services/user_service.dart';
import './add_edit_user_page.dart';
import '../custom_data_table.dart';
import 'dart:ui'; // Pour utiliser BackdropFilter

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
                      showDialog(
                        context: context,
                        builder: (context) {
                          return AddEditUserModal(
                            user: user,
                            onUserSaved: () {
                              _loadUsers(); // Rafraîchir la liste
                            },
                          );
                        },
                      );
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

          return SingleChildScrollView(
            child: CustomDataTable(
              title: 'Utilisateurs',
              columns: columns,
              rows: rows,
              buttonText: 'Ajouter un utilisateur',
              onButtonPressed: () {
                showDialog(
                  context: context,
                  builder: (context) {
                    return AddEditUserModal(
                      user: null,
                      onUserSaved: () {
                        _loadUsers(); // Rafraîchir la liste
                      },
                    );
                  },
                );
              },
            ),
          );
        },
      ),
    );
  }

  void _confirmDelete(BuildContext context, UserApp user) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return Dialog(
          backgroundColor: Colors
              .transparent, // Rendre le fond transparent pour voir le flou
          child: Stack(
            children: [
              // Appliquer un flou sur l'arrière-plan
              Positioned.fill(
                child: BackdropFilter(
                  filter: ImageFilter.blur(
                      sigmaX: 10.0, sigmaY: 10.0), // Valeur de flou
                  child: Container(
                    color: Colors.black
                        .withOpacity(0), // Pour appliquer un fond transparent
                  ),
                ),
              ),
              // Modal principale
              AlertDialog(
                title: const Text('Supprimer l\'utilisateur'),
                content: Text('Voulez-vous vraiment supprimer ${user.email} ?'),
                actions: [
                  TextButton(
                    onPressed: () => Navigator.of(context).pop(),
                    child: const Text(
                      'Annuler',
                      style: TextStyle(color: Colors.black),
                    ),
                  ),
                  TextButton(
                    onPressed: () async {
                      await UserService.deleteUser(user.id);
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                            content: Text('${user.email} a été supprimé.')),
                      );
                      Navigator.of(context).pop();
                      _loadUsers();
                    },
                    style: TextButton.styleFrom(
                      side: BorderSide(color: Colors.red, width: 2),
                      backgroundColor: Colors.red,
                      padding:
                          EdgeInsets.symmetric(vertical: 12, horizontal: 24),
                    ),
                    child: const Text(
                      'Confirmer',
                      style: TextStyle(color: Colors.white),
                    ),
                  ),
                ],
              ),
            ],
          ),
        );
      },
    );
  }
}
