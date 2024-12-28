import 'package:flutter/material.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/services/user_service.dart';

class AddEditUserPage extends StatefulWidget {
  final UserApp? user;

  const AddEditUserPage({super.key, this.user});

  @override
  _AddEditUserPageState createState() => _AddEditUserPageState();
}

class _AddEditUserPageState extends State<AddEditUserPage> {
  final _formKey = GlobalKey<FormState>();
  late String _name;
  late String _email;
  late UserRole _selectedRole;

  @override
  void initState() {
    super.initState();
    _name = widget.user?.name ?? '';
    _email = widget.user?.email ?? '';
    _selectedRole = widget.user?.roles.isNotEmpty == true ? widget.user!.roles.first : UserRole.user;
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;

    _formKey.currentState!.save();
    final isEditing = widget.user != null;

    try {
      if (isEditing) {
        // Appel PUT pour mettre à jour l'utilisateur
        await UserService.updateUser(
          widget.user!.id,
          {
            'name': _name,
            'email': _email,
            'roles': [_selectedRole.name], // Assurez-vous d'envoyer le rôle correctement
          },
        );
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Utilisateur mis à jour avec succès.')),
        );
      } else {
        // Code pour ajouter un nouvel utilisateur (non inclus ici)
      }

      Navigator.of(context).pop();
    } catch (error) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Erreur: ${error.toString()}'),
          backgroundColor: Theme.of(context).colorScheme.error,
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final isEditing = widget.user != null;

    return Scaffold(
      appBar: AppBar(
        title: Text(isEditing ? 'Modifier utilisateur' : 'Ajouter utilisateur'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: ListView(
            children: [
              TextFormField(
                initialValue: _name,
                decoration: const InputDecoration(labelText: 'Nom'),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer un nom.';
                  }
                  return null;
                },
                onSaved: (value) {
                  _name = value!;
                },
              ),
              const SizedBox(height: 20),
              TextFormField(
                initialValue: _email,
                decoration: const InputDecoration(labelText: 'Email'),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer un email.';
                  }
                  return null;
                },
                onSaved: (value) {
                  _email = value!;
                },
              ),
              const SizedBox(height: 20),
              // Dropdown pour le rôle
              DropdownButtonFormField<UserRole>(
                value: _selectedRole,
                items: UserRole.values
                    .map((role) => DropdownMenuItem(
                          value: role,
                          child: Text(role.name),
                        ))
                    .toList(),
                onChanged: (value) {
                  setState(() {
                    _selectedRole = value!;
                  });
                },
                decoration: const InputDecoration(labelText: 'Rôle'),
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: _submit,
                child: Text(isEditing ? 'Modifier' : 'Ajouter'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
