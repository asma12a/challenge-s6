import 'package:flutter/material.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/services/user_service.dart';

class AddEditUserModal extends StatefulWidget {
  final UserApp? user;
  final VoidCallback?
      onUserSaved; // Callback pour notifier que l'utilisateur a été ajouté ou modifié

  const AddEditUserModal({super.key, this.user, this.onUserSaved});

  @override
  _AddEditUserModalState createState() => _AddEditUserModalState();
}

class _AddEditUserModalState extends State<AddEditUserModal> {
  final _formKey = GlobalKey<FormState>();
  late String _name;
  late String _email;
  late UserRole _selectedRole;

  @override
  void initState() {
    super.initState();
    if (widget.user != null) {
      _name = widget.user!.name;
      _email = widget.user!.email;
      _selectedRole = widget.user!.roles.isNotEmpty
          ? widget.user!.roles.first
          : UserRole.user;
    } else {
      _name = '';
      _email = '';
      _selectedRole = UserRole.user;
    }
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;

    _formKey.currentState!.save();
    final isEditing = widget.user != null;

    try {
      if (isEditing) {
        await UserService.updateUser(
          widget.user!.id,
          {
            'name': _name,
            'email': _email,
            'roles': [_selectedRole.name],
          },
        );
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Utilisateur mis à jour avec succès.')),
        );
      } else {
        await UserService.createUser(
          {
            'name': _name,
            'email': _email,
            'roles': [_selectedRole.name],
          },
        );
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Utilisateur créé avec succès.')),
        );
      }

      // Appeler le callback pour notifier que l'utilisateur a été sauvegardé
      if (widget.onUserSaved != null) {
        widget.onUserSaved!();
      }

      Navigator.of(context).pop(); // Fermer la modal
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

    return AlertDialog(
      title: Text(isEditing ? 'Modifier l\'utilisateur' : 'Nouvel utilisateur'),
      content: Form(
        key: _formKey,
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
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
            ],
          ),
        ),
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.of(context).pop(),
          style: TextButton.styleFrom(
            backgroundColor: Colors.white,
            foregroundColor: Colors.black,
            side: const BorderSide(color: Colors.black),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          ),
          child: const Text('Annuler'),
        ),
        ElevatedButton(
          onPressed: _submit,
          style: ElevatedButton.styleFrom(
            backgroundColor: Colors.green,
            foregroundColor: Colors.white,
          ),
          child: Text(isEditing ? 'Modifier' : 'Ajouter'),
        ),
      ],
    );
  }
}
