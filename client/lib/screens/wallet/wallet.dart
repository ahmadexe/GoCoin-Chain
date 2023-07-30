import 'package:client/blocs/wallet/wallet_bloc.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'widgets/_body.dart';

class WalletScreen extends StatefulWidget {
  const WalletScreen({super.key});

  @override
  State<WalletScreen> createState() => _WalletScreenState();
}

class _WalletScreenState extends State<WalletScreen> {
  @override
  void initState() {
    super.initState();
    final walletBloc = WalletBloc.get(context, false);
    walletBloc.add(const GetWalletAmount());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text(
          'Wallet',
          style: TextStyle(color: Colors.black),
        ),
        backgroundColor: Colors.white,
        elevation: 0,
      ),
      body: BlocBuilder<WalletBloc, WalletState>(
        builder: (context, state) {
          if (state.amount is WalletAmountLoading || state.amount is WalletAmountDefault) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          }
          return const _Body();
        },
      ),
    );
  }
}
