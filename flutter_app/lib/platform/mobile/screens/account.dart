import 'package:animated_toggle_switch/animated_toggle_switch.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/core/services/sport_service.dart';
import 'package:squad_go/core/services/user_service.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/edit_user.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/offline.dart';
import 'package:squad_go/platform/mobile/widgets/home_widgets/my_events.dart';

import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/platform/mobile/widgets/performances.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class AccountScreen extends StatefulWidget {
  const AccountScreen({super.key});

  @override
  State<AccountScreen> createState() => _AccountScreenState();
}

class _AccountScreenState extends State<AccountScreen> {
  final SportService sportService = SportService();

  Future<void> onRefresh() async {}
  int eventsCount = 0;
  List<Sport> userSports = [];
  List<String> labelNotifs = ["Évenements recommandés", "Évenement 1jour avant", "Personne invité inscrite"];

  @override
  void initState() {
    super.initState();
    fetchUserSport();
  }

  Future<void> fetchUserSport() async {
    try {
      List<Sport> sports = await sportService.getUserSports();
      setState(() {
        userSports = sports;
      });
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch user sports: $e');
    }
  }

  Future<void> _updateUserInfo(BuildContext context, String name, String email) async {
    try {
      final authState = context.read<AuthState>();

      // Mettre à jour le backend avec UserService
      final userId = authState.userInfo?.id ?? '';
      if (userId.isEmpty) {
        throw Exception('ID utilisateur introuvable.');
      }

      final response = await UserService.updateUser(userId, {
        'name': name,
        'email': email,
      });

      authState.setUser(UserApp(
        id: userId,
        name: name,
        email: email,
        roles: authState.userInfo?.roles ?? [],
        apiToken: authState.userInfo!.apiToken,
      ));

      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Utilisateur mis à jour avec succès !')),
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur : ${e.toString()}')),
      );
    }
  }

  Future<void> _updateUserPassword(BuildContext context, String password) async {
    try {
      final authState = context.read<AuthState>();

      // Mettre à jour le backend avec UserService
      final userId = authState.userInfo?.id ?? '';
      if (userId.isEmpty) {
        throw Exception('ID utilisateur introuvable.');
      }

      await UserService.updateUser(userId, {'password': password});

      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Mot de passe mis à jour avec succès !')),
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur : ${e.toString()}')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    var isOnline = context.watch<ConnectivityState>().isConnected;

    return Consumer<AuthState>(
      builder: (context, authState, child) {
        final userInfo = authState.userInfo;
        return SafeArea(
          child: Scaffold(
            body: RefreshIndicator(
              edgeOffset: 40,
              onRefresh: onRefresh,
              child: CustomScrollView(
                physics: const AlwaysScrollableScrollPhysics(),
                slivers: [
                  SliverFillRemaining(
                    child: Column(
                      children: [
                        Flexible(
                          flex: 1,
                          child: Container(
                            margin: const EdgeInsets.only(
                              bottom: 16,
                              left: 16,
                              right: 16,
                            ),
                            decoration: BoxDecoration(
                              color: Theme.of(context).colorScheme.primary.withOpacity(0.03),
                              borderRadius: BorderRadius.all(
                                Radius.circular(16),
                              ),
                            ),
                            child: Stack(
                              children: [
                                Padding(
                                  padding: const EdgeInsets.all(4),
                                  child: Center(
                                    child: Column(
                                        mainAxisAlignment: MainAxisAlignment.center,
                                        crossAxisAlignment: CrossAxisAlignment.center,
                                        children: [
                                          Text(
                                            userInfo?.name ?? "Utilisateur",
                                            style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                                                  fontWeight: FontWeight.bold,
                                                ),
                                          ),
                                          const SizedBox(height: 8),
                                          if (eventsCount > 0) // Vérification si eventsCount est supérieur à 0
                                            Text(
                                              translate?.events_participated(eventsCount) ??
                                                  "A participé à $eventsCount événement${eventsCount > 1 ? 's' : ''}",
                                              style: Theme.of(context).textTheme.bodyMedium,
                                            ),
                                        ]),
                                  ),
                                ),
                                Positioned(
                                    top: 8,
                                    right: 8,
                                    child: IconButton(
                                      icon: Icon(Icons.edit, color: Theme.of(context).colorScheme.primary),
                                      onPressed: () {
                                        if (!isOnline) {
                                          showDialog(
                                            context: context,
                                            builder: (context) => const OfflineDialog(),
                                          );
                                          return;
                                        }
                                        showDialog(
                                          context: context,
                                          builder: (context) {
                                            return EditUserDialog(
                                              onUpdateInfo: (name, email) => _updateUserInfo(context, name, email),
                                              onUpdatePassword: (password) => _updateUserPassword(context, password),
                                            );
                                          },
                                        );
                                      },
                                    ))
                              ],
                            ),
                          ),
                        ),
                        Flexible(
                          flex: 4,
                          child: Container(
                            margin: const EdgeInsets.only(
                              bottom: 16,
                              left: 16,
                              right: 16,
                            ),
                            child: DefaultTabController(
                              length: 3,
                              child: Column(
                                children: [
                                  Container(
                                    height: 40,
                                    decoration: BoxDecoration(
                                      color: Theme.of(context).colorScheme.primary.withOpacity(0.5),
                                      borderRadius: BorderRadius.circular(10),
                                    ),
                                    child: TabBar(
                                      indicatorSize: TabBarIndicatorSize.tab,
                                      dividerColor: Colors.transparent,
                                      indicator: BoxDecoration(
                                        color: Theme.of(context).colorScheme.primary.withOpacity(0.2),
                                        borderRadius: BorderRadius.circular(10),
                                      ),
                                      labelColor: Colors.white,
                                      labelStyle: TextStyle(
                                        fontWeight: FontWeight.bold,
                                      ),
                                      tabs: [
                                        Tab(
                                          child: Text(
                                            translate?.my_events_profile ?? 'Mes events',
                                            overflow: TextOverflow.ellipsis,
                                          ),
                                        ),
                                        Tab(
                                          child: Text(
                                            translate?.performance ?? 'Performances',
                                            overflow: TextOverflow.ellipsis,
                                          ),
                                        ),
                                        Tab(
                                          child: Text(
                                            translate?.settings ?? 'Paramètres',
                                            overflow: TextOverflow.ellipsis,
                                          ),
                                        ),
                                      ],
                                    ),
                                  ),
                                  Flexible(
                                    child: TabBarView(
                                      children: [
                                        Container(
                                          margin: const EdgeInsets.only(top: 16),
                                          decoration: BoxDecoration(
                                            color: Theme.of(context).colorScheme.primary.withOpacity(0.03),
                                            borderRadius: BorderRadius.circular(16),
                                          ),
                                          padding: const EdgeInsets.all(16),
                                          child: Center(
                                            child: Column(
                                              children: [
                                                HomeMyEvents(
                                                  onRefresh: onRefresh,
                                                  isHome: false,
                                                  onEventsCountChanged: (count) {
                                                    setState(() {
                                                      eventsCount = count;
                                                    });
                                                  },
                                                  onDistinctSportsFetched: (sports) {
                                                    setState(() {
                                                      userSports = sports;
                                                    });
                                                    //debugPrint("userSports $userSports");
                                                  },
                                                ),
                                              ],
                                            ),
                                          ),
                                        ),
                                        Container(
                                          margin: const EdgeInsets.only(top: 16),
                                          decoration: BoxDecoration(
                                            color: Theme.of(context).colorScheme.primary.withOpacity(0.03),
                                            borderRadius: BorderRadius.circular(16),
                                          ),
                                          padding: const EdgeInsets.all(16),
                                          child: PerformancesHandle(
                                            sports: userSports,
                                          ),
                                        ),
                                        Container(
                                          margin: const EdgeInsets.only(top: 16),
                                          decoration: BoxDecoration(
                                            color: Theme.of(context).colorScheme.primary.withOpacity(0.03),
                                            borderRadius: BorderRadius.circular(16),
                                          ),
                                          padding: const EdgeInsets.all(16),
                                          child: Column(
                                            crossAxisAlignment: CrossAxisAlignment.start,
                                            children: [
                                              // Header
                                              Row(
                                                children: [
                                                  const Icon(
                                                    Icons.notifications_active_outlined,
                                                    size: 24,
                                                    color: Colors.blue,
                                                  ),
                                                  const SizedBox(width: 8),
                                                  Text(
                                                    'Notifications',
                                                    style: TextStyle(
                                                      fontSize: 18,
                                                      fontWeight: FontWeight.bold,
                                                      color: Theme.of(context).colorScheme.primary,
                                                    ),
                                                  ),
                                                ],
                                              ),
                                              const SizedBox(height: 16),

                                              // Notification items
                                              ListView.builder(
                                                shrinkWrap: true,
                                                physics: const NeverScrollableScrollPhysics(),
                                                itemCount: labelNotifs.length,
                                                itemBuilder: (context, index) {
                                                  final stat = labelNotifs[index];
                                                  return Container(
                                                    margin: const EdgeInsets.symmetric(vertical: 8),
                                                    padding: const EdgeInsets.symmetric(horizontal: 8),
                                                    decoration: BoxDecoration(
                                                      color: Theme.of(context)
                                                          .colorScheme
                                                          .primary
                                                          .withAlpha(20),
                                                      borderRadius: BorderRadius.circular(16),
                                                    ),
                                                    child: ListTile(
                                                      title: Text(
                                                        stat,
                                                        style: const TextStyle(
                                                          fontSize: 14,
                                                          fontWeight: FontWeight.w500,
                                                        ),
                                                      ),
                                                      trailing: AnimatedToggleSwitch<bool>.dual(
                                                        current: true,
                                                        first: false,
                                                        second: true,
                                                        onChanged: (value) => {},
                                                        iconBuilder: (value) => value
                                                            ? const Icon(Icons.check_circle, color: Colors.white)
                                                            : const Icon(Icons.cancel, color: Colors.white),
                                                        height: 30,
                                                        indicatorSize: const Size(40, 25),
                                                      ),
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
                                ],
                              ),
                            ),
                          ),
                        )
                      ],
                    ),
                  )
                ],
              ),
            ),
          ),
        );
      },
    );
  }
}
