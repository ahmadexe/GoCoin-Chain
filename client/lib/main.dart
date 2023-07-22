import 'package:client/blocs/wallet/wallet_bloc.dart';
import 'package:client/screens/transactions/transactions.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

void main(List<String> args) {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(create: (_) => WalletBloc()),
      ],
      child: const MaterialApp(
        home: TransactionsScreen(),
      ),
    );
  }
}
