import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/shared_widgets/sign_in.dart';
import 'package:squad_go/platform/web/screens/users/admin_users_page.dart';
import 'package:squad_go/platform/web/screens/events/admin_events_page.dart';
import 'package:squad_go/platform/web/screens/sports/admin_sports_page.dart';
import 'package:squad_go/platform/web/screens/admin_dashboard_page.dart';
import 'package:squad_go/platform/web/screens/sports_stat_labels/admin_sport_stat_labels_page.dart';

class WebHomeScreen extends StatefulWidget {
  const WebHomeScreen({super.key});

  @override
  _WebHomeScreenState createState() => _WebHomeScreenState();
}

class _WebHomeScreenState extends State<WebHomeScreen> {
  int _selectedIndex = 0;

  // Liste des pages correspondantes aux index
  final List<Widget> _pages = [
    const AdminDashboardPage(),
    const AdminUsersPage(),
    const AdminEventsPage(),
    const AdminSportsPage(),
    const AdminSportStatLabelsPage(),
  ];

  @override
  Widget build(BuildContext context) {
    return Consumer<AuthState>(
      builder: (context, authState, _) {
        return Scaffold(
          appBar: PreferredSize(
            preferredSize: Size.fromHeight(60.0),
            child: AppBar(
              backgroundColor: Theme.of(context).colorScheme.primary,
              elevation: 0,
            ),
          ),
          // Structure principale
          body: Row(
            children: [
              // Affiche la sidebar seulement si l'utilisateur est authentifié
              if (authState.isAuthenticated)
                Container(
                  width: 250, // Largeur de la sidebar
                  color: Colors.teal[50],
                  child: ListView(
                    padding: EdgeInsets.zero,
                    children: <Widget>[
                      // Remplacer le DrawerHeader par une image
                      DrawerHeader(
                        child: Center(
                          child: Image.asset(
                            'assets/images/app_icon.png',
                            width: 200,
                            height: 200,
                            fit: BoxFit.contain,
                          ),
                        ),
                      ),
                      _createDrawerItem(
                        icon: Icons.dashboard,
                        text: 'Tableau de Bord',
                        onTap: () {
                          setState(() {
                            _selectedIndex = 0;
                          });
                        },
                      ),
                      _createDrawerItem(
                        icon: Icons.account_circle,
                        text: 'Utilisateurs',
                        onTap: () {
                          setState(() {
                            _selectedIndex = 1;
                          });
                        },
                      ),
                      _createDrawerItem(
                        icon: Icons.event,
                        text: 'Événements',
                        onTap: () {
                          setState(() {
                            _selectedIndex = 2;
                          });
                        },
                      ),
                      _createDrawerItem(
                        icon: Icons.sports,
                        text: 'Sports',
                        onTap: () {
                          setState(() {
                            _selectedIndex = 3;
                          });
                        },
                      ),
                      _createDrawerItem(
                        icon: Icons.bar_chart,
                        text: 'Statistiques sportives',
                        onTap: () {
                          setState(() {
                            _selectedIndex =
                                4;
                          });
                        },
                      ),

                      _createDrawerItem(
                        icon: Icons.exit_to_app,
                        text: 'Déconnexion',
                        onTap: () {
                          authState.logout();
                        },
                      ),
                    ],
                  ),
                ),
              // Contenu principal (ajusté pour le cas non connecté)
              Expanded(
                child: authState.isAuthenticated
                    ? _pages[_selectedIndex] // Affiche la page sélectionnée
                    : const SignInScreen(),
              ),
            ],
          ),
        );
      },
    );
  }

  Widget _createDrawerItem({
    required IconData icon,
    required String text,
    required GestureTapCallback onTap,
  }) {
    return ListTile(
      // leading: Icon(icon, color: Colors.teal),
      title: Text(
        text,
        style: TextStyle(
          color: Colors.black,
          fontSize: 18,
          fontWeight: FontWeight.w600,
        ),
      ),
      onTap: onTap,
    );
  }
}
