import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/services/user_service.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class AddEditUserModal extends StatefulWidget {
  final UserApp? user;
  final VoidCallback? onUserSaved;

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

      if (widget.onUserSaved != null) {
        widget.onUserSaved!();
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
    final translate = AppLocalizations.of(context);
    final isEditing = widget.user != null;

    return BackdropFilter(
      filter:
          ImageFilter.blur(sigmaX: 10, sigmaY: 10), // Flou de l'arrière-plan
      child: Dialog(
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
        backgroundColor: Colors.white,
        child: Container(
          width: MediaQuery.of(context).size.width *
              0.8, // Limite à 80% de la largeur de l'écran
          constraints: const BoxConstraints(
            maxWidth: 400, // Largeur maximale fixe pour la modale
          ),
          padding: const EdgeInsets.all(16),
          child: Form(
            key: _formKey,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                // Icône de fermeture en haut à droite
                Align(
                  alignment: Alignment.topRight,
                  child: IconButton(
                    icon: const Icon(Icons.close),
                    onPressed: () => Navigator.of(context).pop(),
                    padding: EdgeInsets.zero,
                    color: Colors.black,
                  ),
                ),
                Text(
                  isEditing ? 'Modifier l\'utilisateur' : 'Nouvel utilisateur',
                  textAlign: TextAlign.center,
                  style: const TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                const SizedBox(height: 20),
                TextFormField(
                  initialValue: _name,
                  decoration: InputDecoration(
                    labelText: 'Nom',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
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
                const SizedBox(height: 16),
                TextFormField(
                  initialValue: _email,
                  decoration: InputDecoration(
                    labelText: 'Email',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
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
                const SizedBox(height: 16),
                DropdownButtonFormField<UserRole>(
                  value: _selectedRole,
                  items: UserRole.values
                      .map(
                        (role) => DropdownMenuItem(
                          value: role,
                          child: Text(role.name),
                        ),
                      )
                      .toList(),
                  onChanged: (value) {
                    setState(() {
                      _selectedRole = value!;
                    });
                  },
                  decoration: InputDecoration(
                    labelText: translate?.role ?? 'Rôle',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
                const SizedBox(height: 24),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    TextButton(
                      onPressed: () => Navigator.of(context).pop(),
                      style: TextButton.styleFrom(
                        foregroundColor: const Color.fromARGB(
                            255, 0, 0, 0), // Texte en blanc
                        backgroundColor: Colors.white, // Fond blanc
                        side: BorderSide(color: Colors.black), // Bordure noire
                        padding: const EdgeInsets.symmetric(
                            horizontal: 24, vertical: 12),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      child: const Text('Annuler'),
                    ),
                    ElevatedButton(
                      onPressed: _submit,
                      style: ElevatedButton.styleFrom(
                        foregroundColor: const Color.fromARGB(255, 255, 255, 255), 
                        backgroundColor: Theme.of(context).primaryColor,
                        padding: const EdgeInsets.symmetric(
                            horizontal: 24, vertical: 12),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      child: Text(isEditing ? 'Modifier' : 'Ajouter'),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
