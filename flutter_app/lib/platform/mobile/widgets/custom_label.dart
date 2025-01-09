import 'package:flutter/material.dart';

class CustomLabel extends StatelessWidget {
  final String label;
  final IconData? icon;
  final Color? color;
  final Color? iconColor;
  final Color? backgroundColor;

  const CustomLabel({
    super.key,
    required this.label,
    this.icon,
    this.color,
    this.iconColor,
    this.backgroundColor,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: backgroundColor ?? Colors.grey.shade300,
        borderRadius: BorderRadius.circular(8),
      ),
      padding: EdgeInsets.symmetric(horizontal: 6, vertical: 3),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          if (icon != null)
            Icon(
              icon,
              color: iconColor,
              size: MediaQuery.of(context).size.width * 0.05,
            ),
          if (icon != null)
            SizedBox(width: MediaQuery.of(context).size.width * 0.02),
          Flexible(
            child: Text(
              label,
              style: TextStyle(color: color),
              overflow: TextOverflow.ellipsis,
            ),
          ),
        ],
      ),
    );
  }
}
