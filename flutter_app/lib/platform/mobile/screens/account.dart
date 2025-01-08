import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/services/sport_service.dart';
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




  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    final userInfo = context.read<AuthState>().userInfo;
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
                        child: Padding(
                          padding: const EdgeInsets.all(4),
                          child: Center(
                            child: Column(
                                mainAxisAlignment: MainAxisAlignment.start,
                                crossAxisAlignment: CrossAxisAlignment.center,
                                children: [
                                  CircleAvatar(
                                    radius: 25,
                                    backgroundColor: Colors.grey[300],
                                    child: const Icon(
                                      Icons.person,
                                      size: 25,
                                      color: Colors.grey,
                                    ),
                                  ),
                                  const SizedBox(height: 8),
                                  Text(
                                    userInfo?.name ?? "Utilisateur",
                                    style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                                          fontWeight: FontWeight.bold,
                                        ),
                                  ),
                                  const SizedBox(height: 8),
                                  Text('A participer à $eventsCount events')
                                ]),
                          ),
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
                                        translate?.my_events_profile ?? 'Mes Events',
                                        overflow: TextOverflow.ellipsis,
                                      ),
                                    ),
                                    Tab(
                                      child: Text(
                                        translate?.performance ?? 'Performance',
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
                                        child: Text(translate?.my_events_profile ?? 'Mes events'),
                                      ),
                                    ),
                                    Container(
                                      margin: const EdgeInsets.only(top: 16),
                                      decoration: BoxDecoration(
                                        color: Theme.of(context).colorScheme.primary.withOpacity(0.03),
                                        borderRadius: BorderRadius.circular(16),
                                      ),
                                      padding: const EdgeInsets.all(16),
                                      child: Center(child: Text(translate?.performance ?? 'Performances')),
                                    ),
                                    Container(
                                      margin: const EdgeInsets.only(top: 16),
                                      decoration: BoxDecoration(
                                        color: Theme.of(context).colorScheme.primary.withOpacity(0.03),
                                        borderRadius: BorderRadius.circular(16),
                                      ),
                                      padding: const EdgeInsets.all(16),
                                      child: Center(child: Text(translate?.settings ?? 'Paramètres')),
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
  }
}
