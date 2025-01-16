import 'package:flutter/material.dart';

class CustomDataTable extends StatelessWidget {
  final String title;
  final List<DataColumn> columns;
  final List<DataRow> rows;
  final String buttonText;
  final VoidCallback onButtonPressed;

  const CustomDataTable({
    super.key,
    required this.title,
    required this.columns,
    required this.rows,
    required this.buttonText,
    required this.onButtonPressed,
  });

  @override
  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(16.0),
      padding: const EdgeInsets.all(16.0),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(8),
        border: Border.all(
          color: Colors.grey[300]!,
          width: 1,
        ),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Center(
            child: Text(
              title,
              style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                    fontWeight: FontWeight.bold,
                    color: Theme.of(context).primaryColor,
                  ),
            ),
          ),
          const SizedBox(height: 16),
          Align(
            alignment: Alignment.topRight,
            child: ElevatedButton(
              onPressed: onButtonPressed,
              style: ElevatedButton.styleFrom(
                backgroundColor: Theme.of(context).colorScheme.primary,
                foregroundColor: Colors.white,
                padding:
                    const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
              ),
              child: Text(buttonText),
            ),
          ),
          const SizedBox(height: 16),

          // ✅ Correction : Ajout d'un défilement vertical
          Expanded(
            child: SingleChildScrollView(
              child: LayoutBuilder(
                builder: (context, constraints) {
                  return SingleChildScrollView(
                    scrollDirection: Axis.horizontal,
                    child: SizedBox(
                      width: constraints.maxWidth,
                      child: DataTable(
                        headingRowColor: WidgetStateColor.resolveWith(
                            (states) => Colors.grey[200]!),
                        columns: columns,
                        rows: rows,
                        columnSpacing: 16,
                      ),
                    ),
                  );
                },
              ),
            ),
          ),
        ],
      ),
    );
  }
}
