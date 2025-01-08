import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/platform/mobile/widgets/home_widgets/sport_performances.dart';

import '../../../core/models/sport.dart';
import '../../../core/providers/auth_state_provider.dart';

class PerformancesHandle extends StatefulWidget {
  final List<Sport> sports;

  const PerformancesHandle({super.key, required this.sports});

  @override
  State<PerformancesHandle> createState() => _PerformancesHandleState();
}

class _PerformancesHandleState extends State<PerformancesHandle> {
  bool userHasSport = false;

  @override
  void initState() {
    super.initState();
    userHasSport = widget.sports.isNotEmpty;
  }

  @override
  Widget build(BuildContext context) {
    final currentUser = context.read<AuthState>().userInfo;
    return DefaultTabController(
        length: widget.sports.length,
        child: Column(
          children: [
            Stack(
              children: [
                SizedBox(
                  height: 40,
                  child: TabBar(
                    isScrollable: true,
                    padding: EdgeInsets.only(right: 100),
                    dividerColor: Colors.transparent,
                    indicatorSize: TabBarIndicatorSize.tab,
                    tabAlignment: TabAlignment.start,
                    indicator: BoxDecoration(
                      color: Theme.of(context).colorScheme.secondary.withOpacity(0.5),
                      borderRadius: BorderRadius.circular(10),
                    ),
                    labelColor: Colors.white,
                    labelStyle: TextStyle(fontWeight: FontWeight.bold),
                    labelPadding: EdgeInsets.only(left: 16),
                    tabs: widget.sports.map(
                      (sport) {
                        return Tab(
                          child: Row(
                            children: [
                              Text(sport.name.name[0].toUpperCase() + sport.name.name.substring(1)),
                              SizedBox(width: 16),
                            ],
                          ),
                        );
                      },
                    ).toList(),
                  ),
                )
              ],
            ),
            Flexible(
              child: TabBarView(
                  children: widget.sports.map(
                      (sport) => Column(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          SportPerformances(
                              sportId: sport.id,
                              userId: currentUser!.id,
                          )
                        ],
                      )
                  ).toList(),
              ),
            )
          ],
        ));
  }
}