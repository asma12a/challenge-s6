import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';

class EditUserDialog extends StatefulWidget {
  final Future<void> Function(String name, String email)? onUpdateInfo;
  final Future<void> Function(String password)? onUpdatePassword;

  const EditUserDialog({
    super.key,
    this.onUpdateInfo,
    this.onUpdatePassword,
  });

  @override
  State<EditUserDialog> createState() => _EditUserDialogState();
}

class _EditUserDialogState extends State<EditUserDialog> {
  final _formKeyInfo = GlobalKey<FormState>();
  final _formKeyPassword = GlobalKey<FormState>();

  late TextEditingController nameController;
  late TextEditingController emailController;
  final TextEditingController passwordController = TextEditingController();

  @override
  void initState() {
    super.initState();

    final userInfo = context.read<AuthState>().userInfo;

    nameController = TextEditingController(text: userInfo?.name ?? '');
    emailController = TextEditingController(text: userInfo?.email ?? '');
  }

  @override
  void dispose() {
    nameController.dispose();
    emailController.dispose();
    passwordController.dispose();
    super.dispose();
  }

  void _updateUserInfo() async {
    if (_formKeyInfo.currentState!.validate()) {
      final name = nameController.text.trim();
      final email = emailController.text.trim();

      if (widget.onUpdateInfo != null) {
        try {
          // Appeler la fonction de mise à jour fournie
          await widget.onUpdateInfo!(name, email);

          // Mettre à jour le contexte `AuthState`
          final authState = context.read<AuthState>();
          authState.setUser(UserApp(
            id: authState.userInfo?.id ?? '',
            name: name,
            email: email,
            roles: authState.userInfo?.roles ?? [],
            apiToken: authState.userInfo!.apiToken,
          ));
          Navigator.of(context).pop();

          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
                content: Text('Informations utilisateur mises à jour !')),
          );
        } catch (e) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Erreur : ${e.toString()}')),
          );
        }
      }
    }
  }

  void _updateUserPassword() async {
    if (_formKeyPassword.currentState!.validate()) {
      final password = passwordController.text.trim();

      if (widget.onUpdatePassword != null) {
        try {
          await widget.onUpdatePassword!(password);

          // Afficher un message de succès
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('Mot de passe mis à jour !')),
          );
        } catch (e) {
          // Afficher un message d'erreur
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Erreur : ${e.toString()}')),
          );
        }
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: SingleChildScrollView(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                'Modifier les informations utilisateur',
                style: const TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 16),

              // Formulaire pour les informations utilisateur
              Form(
                key: _formKeyInfo,
                child: Column(
                  children: [
                    TextFormField(
                      controller: nameController,
                      decoration: const InputDecoration(
                        labelText: 'Nom',
                        border: OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.trim().isEmpty) {
                          return 'Veuillez entrer un nom.';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16),
                    TextFormField(
                      controller: emailController,
                      decoration: const InputDecoration(
                        labelText: 'Email',
                        border: OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.trim().isEmpty) {
                          return 'Veuillez entrer une adresse email.';
                        }
                        if (!RegExp(r'^[^@]+@[^@]+\.[^@]+').hasMatch(value)) {
                          return 'Veuillez entrer une adresse email valide.';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16),
                    ElevatedButton.icon(
                      onPressed: _updateUserInfo,
                      icon: const Icon(Icons.save),
                      label: const Text('Sauvegarder'),
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.blue,
                        foregroundColor: Colors.white,
                      ),
                    ),
                  ],
                ),
              ),

              const SizedBox(height: 32),

              Text(
                'Changer le mot de passe',
                style: const TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 16),

              // Formulaire pour le mot de passe
              Form(
                key: _formKeyPassword,
                child: Column(
                  children: [
                    TextFormField(
                      controller: passwordController,
                      obscureText: true,
                      decoration: const InputDecoration(
                        labelText: 'Nouveau mot de passe',
                        border: OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.trim().isEmpty) {
                          return 'Veuillez entrer un mot de passe.';
                        }
                        if (value.length < 6) {
                          return 'Le mot de passe doit contenir au moins 6 caractères.';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16),
                    ElevatedButton.icon(
                      onPressed: _updateUserPassword,
                      icon: const Icon(Icons.lock),
                      label: const Text('Changer mot de passe'),
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.green,
                        foregroundColor: Colors.white,
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
