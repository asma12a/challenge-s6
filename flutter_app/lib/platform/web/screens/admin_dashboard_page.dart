import 'package:flutter/material.dart';
import 'package:fl_chart/fl_chart.dart';
import 'package:squad_go/core/services/user_service.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/services/sport_service.dart';
import 'package:squad_go/core/models/user_app.dart';

class AdminDashboardPage extends StatelessWidget {
  const AdminDashboardPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[100],
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Titre de bienvenue
              Center(
                child: Text(
                  'Bienvenue dans SquadGO 🎉',
                  style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                        fontWeight: FontWeight.bold,
                        color: Colors.teal,
                      ),
                ),
              ),
              const SizedBox(height: 20),

              // Section des cartes de résumé
              FutureBuilder<List<UserApp>>(
                future: UserService.getUsers(),
                builder: (context, snapshot) {
                  if (snapshot.connectionState == ConnectionState.waiting) {
                    return const Center(child: CircularProgressIndicator());
                  } else if (snapshot.hasError) {
                    return Center(child: Text('Erreur: ${snapshot.error}'));
                  } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
                    return const Center(
                        child: Text('Aucun utilisateur trouvé.'));
                  } else {
                    final List<UserApp> users = snapshot.data!;
                    final userCount = users.length;

                    return Row(
                      mainAxisAlignment: MainAxisAlignment.spaceAround,
                      children: [
                        _buildSummaryCard(
                          context,
                          icon: Icons.account_circle,
                          label: 'Utilisateurs',
                          value: userCount,
                          color: Colors.teal,
                        ),
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
                              color: Colors.orange,
                            );
                          },
                        ),
                        FutureBuilder<int>(
                          future: SportService.getSports()
                              .then((sports) => sports.length),
                          builder: (context, sportSnapshot) {
                            return _buildSummaryCard(
                              context,
                              icon: Icons.sports,
                              label: 'Sports',
                              value: sportSnapshot.data ?? 0,
                              color: Colors.blue,
                            );
                          },
                        ),
                      ],
                    );
                  }
                },
              ),

              const SizedBox(height: 30),

              // Section des graphiques
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: [
                  _buildPieChartWithFuture(
                    context,
                    title: 'Utilisateurs',
                    futureData: UserService.getUsers(),
                    processData: (users) {
                      final adminCount = users
                          .where((user) => user.roles.contains(UserRole.admin))
                          .length;
                      final userCount = users.length - adminCount;
                      return {'Admins': adminCount, 'Simples': userCount};
                    },
                  ),
                  _buildPieChartWithFuture(
                    context,
                    title: 'Événements',
                    futureData: EventService().getEvents(),
                    processData: (events) {
                      final publicEvents = events
                          .where((event) => event['type'] == 'public')
                          .length;
                      final privateEvents = events.length - publicEvents;
                      return {'Publics': publicEvents, 'Privés': privateEvents};
                    },
                  ),
                ],
              ),
            ],
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
    return Card(
      elevation: 6,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(16),
      ),
      child: Container(
        width: 110,
        height: 130,
        padding: const EdgeInsets.all(12),
        decoration: BoxDecoration(
          color: color.withOpacity(0.9),
          borderRadius: BorderRadius.circular(16),
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(icon, size: 36, color: Colors.white),
            const SizedBox(height: 10),
            Text(
              value.toString(),
              style: const TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
                color: Colors.white,
              ),
            ),
            const SizedBox(height: 4),
            Text(
              label,
              style: const TextStyle(fontSize: 14, color: Colors.white70),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildPieChartWithFuture<T>(
    BuildContext context, {
    required String title,
    required Future<List<T>> futureData,
    required Map<String, int> Function(List<T>) processData,
  }) {
    return FutureBuilder<List<T>>(
      future: futureData,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const SizedBox(
            height: 200,
            width: 200,
            child: Center(child: CircularProgressIndicator()),
          );
        } else if (snapshot.hasError) {
          return Text('Erreur: ${snapshot.error}');
        } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
          return const Text('Aucune donnée disponible');
        } else {
          final data = processData(snapshot.data!);
          return _buildPieChart(title: title, data: data);
        }
      },
    );
  }

  Widget _buildPieChart({
    required String title,
    required Map<String, int> data,
  }) {
    final total = data.values.reduce((a, b) => a + b);

    return Column(
      children: [
        Text(
          title,
          style: const TextStyle(
            fontSize: 18,
            fontWeight: FontWeight.bold,
            color: Colors.teal,
          ),
        ),
        const SizedBox(height: 20),
        SizedBox(
          height: 200,
          width: 200,
          child: PieChart(
            PieChartData(
              sections: data.entries.map((entry) {
                final percentage = (entry.value / total) * 100;
                return PieChartSectionData(
                  value: percentage,
                  color: entry.key == 'Admins' || entry.key == 'Publics'
                      ? Colors.teal
                      : Colors.orange,
                  radius: 50,
                  title: '${entry.key}\n${percentage.toStringAsFixed(1)}%',
                  titleStyle: const TextStyle(
                    fontSize: 14,
                    fontWeight: FontWeight.bold,
                    color: Colors.white,
                  ),
                );
              }).toList(),
              sectionsSpace: 2,
            ),
          ),
        ),
      ],
    );
  }
}
