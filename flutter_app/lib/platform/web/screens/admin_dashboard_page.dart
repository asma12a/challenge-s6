import 'package:flutter/material.dart';
import 'package:squad_go/core/services/user_service.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/services/sport_service.dart';
import 'package:squad_go/core/models/user_app.dart';

class AdminDashboardPage extends StatelessWidget {
  const AdminDashboardPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[50],
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center, 
              children: [
                Padding(
                  padding: const EdgeInsets.only(top: 40.0),
                  child: Center(
                    child: Text(
                      'Bienvenue dans SquadGO 🎉',
                      style: Theme.of(context).textTheme.titleLarge?.copyWith(
                            fontWeight: FontWeight.bold,
                            color: Theme.of(context).primaryColor,
                          ),
                    ),
                  ),
                ),
                const SizedBox(height: 20),

                // Section des cartes de résumé centrées
                FutureBuilder<List<UserApp>>(
                  future: UserService.getUsers(),
                  builder: (context, snapshot) {
                    if (snapshot.connectionState == ConnectionState.waiting) {
                      return const Center(child: CircularProgressIndicator());
                    } else if (snapshot.hasError) {
                      return Center(child: Text('Erreur: ${snapshot.error}'));
                    } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
                      return const Center(child: Text('Aucun utilisateur trouvé.'));
                    } else {
                      final List<UserApp> users = snapshot.data!;
                      final userCount = users.length;

                      return Padding(
                        padding: const EdgeInsets.symmetric(vertical: 30.0),
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.center, // Centre les cartes
                          children: [
                            _buildSummaryCard(
                              context,
                              icon: Icons.account_circle,
                              label: 'Utilisateurs',
                              value: userCount,
                              color: Theme.of(context).primaryColor,
                            ),
                            const SizedBox(width: 20),
                            FutureBuilder<int>(
                              future: EventService()
                                  .getEvents()
                                  .then((events) => events.length),
                              builder: (context, eventSnapshot) {
                                return _buildSummaryCard(
                                  context,
                                  icon: Icons.event,
                                  label: 'Événements',
                                  value: eventSnapshot.data ?? 0,
                                  color: Theme.of(context).primaryColorLight, // Teal ici
                                );
                              },
                            ),
                            const SizedBox(width: 20),
                            FutureBuilder<int>(
                              future: SportService.getSports()
                                  .then((sports) => sports.length),
                              builder: (context, sportSnapshot) {
                                return _buildSummaryCard(
                                  context,
                                  icon: Icons.sports,
                                  label: 'Sports',
                                  value: sportSnapshot.data ?? 0,
                                  color: Theme.of(context).primaryColor,
                                );
                              },
                            ),
                          ],
                        ),
                      );
                    }
                  },
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildSummaryCard(
    BuildContext context, {
    required IconData icon,
    required String label,
    required int value,
    required Color color,
  }) {
    return Expanded(
      child: Card(
        elevation: 6,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(16),
        ),
        child: Container(
          padding: const EdgeInsets.all(16),
          decoration: BoxDecoration(
            color: color.withOpacity(0.9),
            borderRadius: BorderRadius.circular(16),
          ),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(icon, size: 40, color: Colors.white),
              const SizedBox(height: 10),
              Text(
                value.toString(),
                style: const TextStyle(
                  fontSize: 26,
                  fontWeight: FontWeight.bold,
                  color: Colors.white,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                label,
                style: const TextStyle(fontSize: 16, color: Colors.white70),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
