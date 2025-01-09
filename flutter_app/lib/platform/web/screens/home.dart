import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/shared_widgets/sign_in.dart';
import 'package:squad_go/platform/web/screens/users/admin_users_page.dart';
import 'package:squad_go/platform/web/screens/events/admin_events_page.dart';
import 'package:squad_go/platform/web/screens/sports/admin_sports_page.dart';
import 'package:squad_go/platform/web/screens/admin_dashboard_page.dart';
import 'package:squad_go/platform/web/screens/sports_stat_labels/admin_sport_stat_labels_page.dart';
import 'package:squad_go/platform/web/screens/logs/admin_logs_page.dart';

class WebHomeScreen extends StatefulWidget {
  const WebHomeScreen({super.key});

  @override
  _WebHomeScreenState createState() => _WebHomeScreenState();
}

class _WebHomeScreenState extends State<WebHomeScreen> {
  int _selectedIndex = 0;
  String? _errorMessage;

  final List<Widget> _pages = [
    const AdminDashboardPage(),
    const AdminUsersPage(),
    const AdminEventsPage(),
    const AdminSportsPage(),
    const AdminSportStatLabelsPage(),
        const AdminLogsPage(),

  ];

  @override
  Widget build(BuildContext context) {
    return Consumer<AuthState>(
      builder: (context, authState, _) {
        if (authState.isAuthenticated && !authState.isAdmin) {
          WidgetsBinding.instance.addPostFrameCallback((_) {
            setState(() {
              _errorMessage = "Droits insuffisants. Accès refusé.";
            });
            authState.logout();
          });
        }

        return Scaffold(
          appBar: authState.isAuthenticated && authState.isAdmin
              ? null
              : AppBar(
                  backgroundColor: Theme.of(context).colorScheme.primary,
                  elevation: 0,
                ),
          body: Row(
            children: [
              if (authState.isAuthenticated && authState.isAdmin)
                _buildSidebar(),
              Expanded(
                child: authState.isAuthenticated && authState.isAdmin
                    ? _pages[_selectedIndex]
                    : _buildSignInOrError(authState),
              ),
            ],
          ),
        );
      },
    );
  }

  Widget _buildSignInOrError(AuthState authState) {
    if (_errorMessage != null) {
      return Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(
              Icons.warning_amber_rounded,
              size: 100,
              color: Colors.redAccent,
            ),
            const SizedBox(height: 20),
            Text(
              _errorMessage!,
              style: TextStyle(
                color: Colors.grey[800],
                fontSize: 20,
                fontWeight: FontWeight.bold,
                fontFamily: 'Roboto',
              ),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 30),
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                backgroundColor: Theme.of(context).primaryColor, 
                foregroundColor: Colors.white,
                padding: const EdgeInsets.symmetric(
                  horizontal: 40,
                  vertical: 15,
                ),
                textStyle: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                ),
              ),
              onPressed: () {
                setState(() {
                  _errorMessage = null;
                });
              },
              child: const Text('Retour à la page de connexion'),
            ),
          ],
        ),
      );
    }
    return const SignInScreen();
  }

  Widget _buildSidebar() {
    return Container(
      width: 250,
      color: Colors.teal[50],
      child: Column(
        children: <Widget>[
          DrawerHeader(
            child: Center(
              child: Image.asset(
                'assets/images/app_icon.png',
                width: 120,
                height: 120,
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
                _selectedIndex = 4;
              });
            },
          ),
          _createDrawerItem(
            icon: Icons.event,
            text: 'Logs',
            onTap: () {
              setState(() {
                _selectedIndex = 5;
              });
            },
          ),
          const Divider(),
          Spacer(),
          _createDrawerItem(
            icon: Icons.exit_to_app,
            text: 'Déconnexion',
            onTap: () {
              Provider.of<AuthState>(context, listen: false).logout();
            },
          ),
        ],
      ),
    );
  }

  Widget _createDrawerItem({
    required IconData icon,
    required String text,
    required GestureTapCallback onTap,
  }) {
    return ListTile(
      leading: Icon(icon, color: Colors.teal),
      title: Text(
        text,
        style: const TextStyle(
          color: Colors.black,
          fontSize: 16,
          fontFamily: 'Poppins',
          fontWeight: FontWeight.w600,
        ),
      ),
      onTap: onTap,
    );
  }
}
