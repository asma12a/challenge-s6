import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

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
  bool _isPasswordVisible = false;
  String? passwordErrorMessage;

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
    final translate = AppLocalizations.of(context);
    if (_formKeyInfo.currentState!.validate()) {
      final name = nameController.text.trim();
      final email = emailController.text.trim();

      if (widget.onUpdateInfo != null) {
        try {
          await widget.onUpdateInfo!(name, email);
          Navigator.of(context).pop();

        } catch (e) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(
                '${translate?.error ?? "Erreur:"} ${e.toString()}',
              ),
            ),
          );
        }
      }
    }
  }

  void _updateUserPassword() async {
    final translate = AppLocalizations.of(context);
    if (_formKeyPassword.currentState!.validate()) {
      final password = passwordController.text.trim();

      if (widget.onUpdatePassword != null) {
        try {
          await widget.onUpdatePassword!(password);
          setState(() {
            passwordErrorMessage = null;
          });
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text(translate?.updated_password ?? 'Mot de passe mis à jour !')),
          );
        } catch (e) {
          debugPrint("exception $e");
          setState(() {
            passwordErrorMessage = e.toString();
          });
        }
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: SingleChildScrollView(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(translate?.edit_user_infos ??
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
                      decoration: InputDecoration(
                        labelText: translate?.name ?? 'Nom',
                        border: const OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.trim().isEmpty) {
                          return translate?.empty_user_name ?? 'Veuillez entrer un nom.';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16),
                    TextFormField(
                      controller: emailController,
                      decoration: InputDecoration(
                        labelText: translate?.email_label ?? 'Email',
                        border: const OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.trim().isEmpty) {
                          return translate?.empty_email ?? 'Veuillez entrer une adresse email.';
                        }
                        if (!RegExp(r'^[^@]+@[^@]+\.[^@]+').hasMatch(value)) {
                          return translate?.valid_email ?? 'Veuillez entrer une adresse email valide.';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16),
                    ElevatedButton.icon(
                      onPressed: _updateUserInfo,
                      icon: const Icon(Icons.save),
                      label: Text(translate?.save_event ?? 'Sauvegarder'),
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.blue,
                        foregroundColor: Colors.white,
                      ),
                    ),
                  ],
                ),
              ),

              const SizedBox(height: 32),

              Text(translate?.change_password ??
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
                      obscureText: !_isPasswordVisible,
                      onTapOutside: (event) => FocusScope.of(context).unfocus(),
                      decoration: InputDecoration(
                        labelText: translate?.new_password ?? 'Nouveau mot de passe',
                        border: const OutlineInputBorder(),
                        suffixIcon: IconButton(
                          icon: Icon(_isPasswordVisible ? Icons.visibility_off : Icons.visibility),
                          onPressed: () {
                            setState(() {
                              _isPasswordVisible = !_isPasswordVisible;
                            });
                          },
                        ),
                      ),
                      validator: (value) {
                        if (value == null || value.trim().isEmpty) {
                          return translate?.empty_password ?? 'Veuillez entrer un mot de passe.';
                        }
                        if (value.length < 6) {
                          return translate?.six_char ?? 'Le mot de passe doit contenir au moins 6 caractères.';
                        }
                        return null;
                      },
                    ),
                    if (passwordErrorMessage != null)
                      Padding(
                        padding: const EdgeInsets.only(top: 8.0),
                        child: Text(
                          passwordErrorMessage!,
                          style: TextStyle(
                            color: Colors.red,
                            fontSize: 14,
                          ),
                        ),
                      ),
                    const SizedBox(height: 16),
                    ElevatedButton.icon(
                      onPressed: _updateUserPassword,
                      icon: const Icon(Icons.lock),
                      label: Text(translate?.change_password ?? 'Changer mot de passe'),
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
