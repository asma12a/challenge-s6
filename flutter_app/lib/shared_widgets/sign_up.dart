import 'package:flutter/material.dart';
import 'package:squad_go/core/services/auth_service.dart';
import 'package:squad_go/shared_widgets/sign_in.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

import '../platform/mobile/widgets/logo.dart';

class SignUpScreen extends StatelessWidget {
  const SignUpScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final bool isSmallScreen = MediaQuery.of(context).size.width < 600;

    return Scaffold(
        body: Center(
            child: isSmallScreen
                ? SingleChildScrollView(
                    child: const Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        _Title(),
                        _FormContent(),
                        SizedBox(height: 20),
                      ],
                    ),
                  )
                : Container(
                    padding: const EdgeInsets.all(32.0),
                    constraints: const BoxConstraints(maxWidth: 800),
                    child: const Row(
                      children: [
                        Expanded(child: _Title()),
                        Expanded(
                          child: Center(child: _FormContent()),
                        ),
                      ],
                    ),
                  )));
  }
}

class _Title extends StatelessWidget {
  const _Title();

  @override
  Widget build(BuildContext context) {
    // final bool isSmallScreen = MediaQuery.of(context).size.width < 600;
    final translate = AppLocalizations.of(context);
    return Container(
      constraints: const BoxConstraints(maxWidth: 300),
      // Largeur maximale pour alignement avec le formulaire
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          _Logo(),
          Text(
              translate?.join_community ??
                  'Rejoignez votre communauté dès aujourd’hui !',
              style: Theme.of(context).textTheme.titleLarge!.copyWith(
                    color: Colors.black,
                    fontSize: 30,
                  )),
          const SizedBox(height: 50),
        ],
      ),
    );
  }
}

class _FormContent extends StatefulWidget {
  const _FormContent();

  @override
  State<_FormContent> createState() => __FormContentState();
}

class __FormContentState extends State<_FormContent> {
  final authService = AuthService();
  bool _isPasswordVisible = false;

  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  var _enteredPseudo = '';
  var _enteredEmail = '';
  var _enteredPassword = '';

  void _signUp() async {
    final translate = AppLocalizations.of(context);
    try {
      if (_formKey.currentState!.validate()) {
        _formKey.currentState!.save();
        final result = await authService.signUp(
          {
            "name": _enteredPseudo,
            "email": _enteredEmail,
            "password": _enteredPassword,
          },
        );

        if (result?['status'] != null) {
          ScaffoldMessenger.of(context).clearSnackBars();

          switch (result?['status']) {
            case 'error_not_strong_password':
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  backgroundColor: Theme.of(context).colorScheme.errorContainer,
                  content: Text(
                    translate?.not_strong ?? result?['error'],
                    style: TextStyle(
                      color: Theme.of(context).colorScheme.onErrorContainer,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ),
              );
              break;

            case 'error_conflict':
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  backgroundColor: Theme.of(context).colorScheme.errorContainer,
                  content: Text(
                    translate?.error_conflict ?? result?['error'],
                    style: TextStyle(
                      color: Theme.of(context).colorScheme.onErrorContainer,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ),
              );
              break;
          }
        } else {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              duration: const Duration(milliseconds: 5000),
              backgroundColor: Theme.of(context).colorScheme.primary,
              content: Text(
                translate?.email_sent ??
                    'Un email de vérification vous a été envoyé.',
                textAlign: TextAlign.center,
              ),
            ),
          );
          Navigator.of(context).pop();
        }
      }
    } catch (e) {
      ScaffoldMessenger.of(context).clearSnackBars();
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          backgroundColor: Theme.of(context).colorScheme.errorContainer,
          content: Text(
            style: TextStyle(
                color: Theme.of(context).colorScheme.onErrorContainer),
            translate?.error_occurred ??
                'Une erreur est survenue, veuillez réessayer.',
            textAlign: TextAlign.center,
          ),
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return Container(
      constraints: const BoxConstraints(maxWidth: 300),
      child: Form(
        key: _formKey,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please enter some text';
                }
                return null;
              },
              onTapOutside: (event) => FocusScope.of(context).unfocus(),
              style: Theme.of(context).textTheme.titleMedium!.copyWith(
                    color: Theme.of(context).colorScheme.onSurface,
                  ),
              decoration: const InputDecoration(
                labelText: 'Pseudo',
                hintText: 'John Doe',
                prefixIcon: Icon(Icons.person),
                border: OutlineInputBorder(),
              ),
              onSaved: (value) {
                _enteredPseudo = value!;
              },
            ),
            _gap(),
            TextFormField(
              validator: (value) {
                // add email validation
                if (value == null || value.isEmpty) {
                  return 'Please enter some text';
                }

                bool emailValid = RegExp(
                        r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+")
                    .hasMatch(value);
                if (!emailValid) {
                  return 'Please enter a valid email';
                }

                return null;
              },
              onTapOutside: (event) => FocusScope.of(context).unfocus(),
              keyboardType: TextInputType.emailAddress,
              style: Theme.of(context).textTheme.titleMedium!.copyWith(
                    color: Theme.of(context).colorScheme.onSurface,
                  ),
              decoration: InputDecoration(
                labelText: translate?.email_label ?? 'Email',
                hintText: translate?.email_placeholder ??
                    'Entrez votre adresse email',
                prefixIcon: const Icon(Icons.email_outlined),
                border: const OutlineInputBorder(),
              ),
              onSaved: (value) {
                _enteredEmail = value!;
              },
            ),
            _gap(),
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please enter some text';
                }

                if (value.length < 4) {
                  return 'Password must be at least 6 characters';
                }
                return null;
              },
              style: Theme.of(context).textTheme.titleMedium!.copyWith(
                    color: Theme.of(context).colorScheme.onSurface,
                  ),
              autofillHints: const [AutofillHints.newPassword],
              obscureText: !_isPasswordVisible,
              onTapOutside: (event) => FocusScope.of(context).unfocus(),
              decoration: InputDecoration(
                labelText: translate?.password ?? 'Mot de passe',
                hintText: translate?.password_placeholder ??
                    'Entrez votre mot de passe',
                prefixIcon: const Icon(Icons.lock_outline_rounded),
                border: const OutlineInputBorder(),
                suffixIcon: IconButton(
                  icon: Icon(_isPasswordVisible
                      ? Icons.visibility_off
                      : Icons.visibility),
                  onPressed: () {
                    setState(() {
                      _isPasswordVisible = !_isPasswordVisible;
                    });
                  },
                ),
              ),
              onSaved: (value) {
                _enteredPassword = value!;
              },
            ),
            _gap(),
            SizedBox(
              width: double.infinity,
              child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                  shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(4)),
                ),
                onPressed: _signUp,
                child: Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: Text(
                    translate?.signup_button ?? 'Créez votre compte',
                    style: const TextStyle(
                        fontSize: 16, fontWeight: FontWeight.bold),
                  ),
                ),
              ),
            ),
            const SizedBox(
              height: 20,
            ),
            Align(
              alignment: Alignment.centerLeft,
              child: Row(
                children: [
                  Text(
                    translate?.already_account ?? 'Vous avez déjà un compte ?',
                    style: const TextStyle(color: Colors.black),
                  ),
                  const SizedBox(
                    width: 10,
                  ),
                  InkWell(
                    child: Text(
                      translate?.login_button ?? 'Connectez-vous',
                      style: TextStyle(
                        color: Theme.of(context).colorScheme.primary,
                      ),
                    ),
                    onTap: () {
                      Navigator.of(context).pushReplacement(
                        MaterialPageRoute(
                          builder: (ctx) => const SignInScreen(),
                        ),
                      );
                    },
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _gap() => const SizedBox(height: 16);
}

class _Logo extends StatelessWidget {
  const _Logo();

  @override
  Widget build(BuildContext context) {
    final bool isSmallScreen = MediaQuery.of(context).size.width < 600;

    return Column(
      mainAxisSize: MainAxisSize.min,
      children: [
        Logo(
          width: isSmallScreen ? 200 : 300,
        ),
        SizedBox(
          height: 25,
        )
      ],
    );
  }
}
