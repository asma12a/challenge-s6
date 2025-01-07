import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/platform/mobile/widgets/home_widgets/my_events.dart';
import 'package:squad_go/platform/mobile/widgets/home_widgets/recommended_events.dart';

class HomeScreen extends StatefulWidget {
  final bool? shouldRefresh;
  const HomeScreen({super.key, this.shouldRefresh});

  @override
  State<HomeScreen> createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen>
    with SingleTickerProviderStateMixin {
  late AnimationController _animationController;
  final GlobalKey<HomeRecommendedEventsState> recommendedEventsKey =
      GlobalKey<HomeRecommendedEventsState>();
  final GlobalKey<HomeMyEventsState> myEventsKey =
      GlobalKey<HomeMyEventsState>();

  @override
  void initState() {
    super.initState();

    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
      lowerBound: 0,
      upperBound: 1,
    );

    _animationController.forward();

    if (widget.shouldRefresh == true) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        onRefresh();
      });
    }
  }

  @override
  void dispose() {
    _animationController.dispose();
    super.dispose();
  }

  Future<void> onRefresh() async {
    await myEventsKey.currentState?.fetchMyEvents();
    await recommendedEventsKey.currentState?.fetchRecommendedEvents();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      builder: (context, child) => SlideTransition(
        position: Tween(
          begin: const Offset(0, 0.3),
          end: const Offset(0, 0),
        ).animate(
          CurvedAnimation(
            parent: _animationController,
            curve: Curves.easeInOut,
          ),
        ),
        child: child,
      ),
      animation: _animationController,
      child: RefreshIndicator(
        onRefresh: onRefresh,
        child: SingleChildScrollView(
          physics: const AlwaysScrollableScrollPhysics(),
          child: Column(
            children: [
              const SizedBox(
                height: 10,
              ),
              HomeMyEvents(
                key: myEventsKey,
                onRefresh: onRefresh,
              ),
              HomeRecommendedEvents(
                key: recommendedEventsKey,
                onRefresh: onRefresh,
              )
            ],
          ),
        ),
      ),
    );
  }
}
