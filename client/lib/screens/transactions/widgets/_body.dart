part of '../transactions.dart';

class _Body extends StatelessWidget {
  const _Body();

  @override
  Widget build(BuildContext context) {
    final appMedia = MediaQuery.sizeOf(context);
    
    return SingleChildScrollView(
      padding: EdgeInsets.all(appMedia.width * 0.05),
      child: const Column(
        children: [
          Text(
            'Transactions',
            style: TextStyle(
              fontSize: 24,
              fontWeight: FontWeight.w600,
            ),
          ),
        ],
      ),
    );
  }
}
