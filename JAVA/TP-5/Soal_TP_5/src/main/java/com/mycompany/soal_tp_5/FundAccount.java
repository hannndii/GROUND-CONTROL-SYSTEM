/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.soal_tp_5;
/**
 *
 * @author tupog
 */
class FundAccount extends Account {
    private double earningRate;

    // Constructor
    public FundAccount(String number, String name, int balance, double earningRate) {
        super(number, name, balance);
        this.earningRate = earningRate;
    }

    @Override
    public void openAccount() {
        System.out.println("Buka akun.");
        System.out.println("Nomor rekening: " + getNumber());
        System.out.println("Pemilik akun: " + getName());
        System.out.println("Saldo: " + "Rp." + getBalance());
        System.out.println("Tingkat pengembalian: " + earningRate + "%");
        System.out.println("Keuntungan telah diperoleh.");
    }
}