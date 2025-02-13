import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/core/utils/tools.dart';
import 'package:squad_go/platform/mobile/widgets/custom_label.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/join_event.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/offline.dart';

class EventCard extends StatelessWidget {
  final Event event;
  final bool hasJoinedEvent;
  final Future<void> Function()? onRefresh;

  const EventCard({
    super.key,
    required this.event,
    this.hasJoinedEvent = false,
    this.onRefresh,
  });

  void onCardClick(BuildContext context) {
    final isOnline = context.read<ConnectivityState>().isConnected;
    if (hasJoinedEvent) {
      if (event.id == null) return;

      context.go("/event/${event.id}", extra: event);
    } else {
      // Show the join event dialog/modal
      if (event.id == null) return;

      if (!isOnline) {
        showDialog(
          context: context,
          builder: (context) => const OfflineDialog(),
        );
        return;
      }

      showDialog(
        context: context,
        builder: (BuildContext context) {
          return JoinEventDialog(
            eventId: event.id!,
            onRefresh: onRefresh,
          );
        },
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return Center(
      child: Card(
        clipBehavior: Clip.hardEdge,
        color: Theme.of(context).colorScheme.surface,
        child: InkWell(
          splashColor: event.sport.color?.withAlpha(30) ??
              Theme.of(context).colorScheme.secondary.withAlpha(30),
          highlightColor: event.sport.color?.withAlpha(30) ??
              Theme.of(context).colorScheme.secondary.withAlpha(30),
          onTap: () {
            onCardClick(context);
          },
          child: SizedBox(
            width: 300,
            height: 200,
            child: Stack(children: [
              Column(
                children: [
                  Expanded(
                    flex: 3,
                    child: event.sport.imageUrl != null
                        ? Image.network(
                            fit: BoxFit.cover,
                            width: double.infinity,
                            event.sport.imageUrl as String,
                          )
                        : Icon(
                            Icons.hide_image,
                            size: 80,
                            color: event.sport.color?.withAlpha(50) ??
                                Theme.of(context)
                                    .colorScheme
                                    .secondary
                                    .withAlpha(50),
                          ),
                  ),
                  Expanded(
                    flex: 4,
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        ListTile(
                          leading: Icon(
                            sportIcon[event.sport.name],
                            size: 32,
                            color: event.sport.color?.withOpacity(0.5),
                          ),
                          title: Text(
                            event.name,
                            maxLines: 1,
                            overflow: TextOverflow.ellipsis,
                          ),
                          subtitle: Row(
                            children: [
                              Expanded(
                                child: Row(
                                  children: [
                                    Icon(
                                      Icons.place,
                                      size: 16,
                                    ),
                                    SizedBox(width: 8),
                                    Expanded(
                                      child: Text(
                                        event.address,
                                        overflow: TextOverflow.ellipsis,
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                            ],
                          ),
                        ),
                        Expanded(
                          child: Padding(
                            padding: const EdgeInsets.only(
                              left: 16,
                              right: 16,
                            ),
                            child: Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                CustomLabel(
                                  label: DateFormat('yyyy/MM/dd HH:mm').format(
                                    DateTime.parse(event.date).add(
                                      Duration(hours: 1),
                                    ),
                                  ),
                                  icon: Icons.date_range,
                                  color: getColorBasedOnDate(event.date),
                                  iconColor: getColorBasedOnDate(event.date),
                                  backgroundColor:
                                      getColorBasedOnDate(event.date)
                                          .withAlpha(20),
                                ),
                                TextButton(
                                  onPressed: () {
                                    onCardClick(context);
                                  },
                                  child: Text(
                                    hasJoinedEvent
                                        ? translate?.see_button ?? "VOIR"
                                        : translate?.join_button ?? 'REJOINDRE',
                                  ),
                                )
                              ],
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
              Positioned(
                  right: 10,
                  top: 10,
                  child: Icon(
                    event.type == EventType.match
                        ? Icons.sports
                        : Icons.fitness_center,
                    size: 32,
                    color: event.type == EventType.match ? Colors.white : null,
                  )),
            ]),
          ),
        ),
      ),
    );
  }
}
