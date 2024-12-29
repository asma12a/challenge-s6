import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/shared_widgets/sign_in.dart';
import 'package:squad_go/platform/web/screens/users/admin_users_page.dart';

class WebHomeScreen extends StatelessWidget {
  const WebHomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<AuthState>(
      builder: (context, authState, _) {
        return Scaffold(
          appBar: PreferredSize(
            preferredSize:
                Size.fromHeight(60.0), // Ajuster la hauteur de l'AppBar
            child: AppBar(
              backgroundColor: Colors.teal, // Couleur de l'AppBar
              elevation:
                  0, // Retirer l'ombre de l'AppBar pour un look plus propre
            ),
          ),
          // Drawer (Sidebar)
          drawer: Drawer(
            child: Container(
              color: Colors.teal[50], // Fond clair de la sidebar
              child: ListView(
                padding: EdgeInsets.zero,
                children: <Widget>[
                  const DrawerHeader(
                    decoration: BoxDecoration(
                      color: Colors.teal, // Fond de l'en-tête de la sidebar
                    ),
                    child: Text(
                      'Menu',
                      style: TextStyle(
                        color: Colors.white,
                        fontSize: 24,
                      ),
                    ),
                  ),
                  _createDrawerItem(
                    icon: Icons.account_circle,
                    text: 'Utilisateurs',
                    onTap: () {
                      // Action de navigation vers le profil utilisateur
                      Navigator.pop(context); // Fermer la sidebar
                    },
                  ),
                  _createDrawerItem(
                    icon: Icons.event,
                    text: 'Événements',
                    onTap: () {
                      // Action de navigation vers la page des événements
                      Navigator.pop(context); // Fermer la sidebar
                    },
                  ),
                  _createDrawerItem(
                    icon: Icons.exit_to_app,
                    text: 'Déconnexion',
                    onTap: () {
                      authState.logout(); // Déconnexion
                      Navigator.pop(context); // Fermer la sidebar
                    },
                  ),
                ],
              ),
            ),
          ),
          body: authState.isAuthenticated
              ? const AdminUsersPage() // Afficher la page admin si l'utilisateur est connecté
              : const SignInScreen(), // Afficher la page de connexion si non authentifié
        );
      },
    );
  }

  // Méthode pour créer un item personnalisé pour le Drawer
  Widget _createDrawerItem(
      {required IconData icon,
      required String text,
      required GestureTapCallback onTap}) {
    return ListTile(
      leading: Icon(icon, color: Colors.teal),
      title: Text(
        text,
        style: TextStyle(
          color: Colors.black, // Couleur du texte
          fontSize: 18,
          fontWeight:
              FontWeight.w600, // Texte en gras pour un meilleur contraste
        ),
      ),
      onTap: onTap,
    );
  }
}
