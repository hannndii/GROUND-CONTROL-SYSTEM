/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.polimorphism;

/**
 *
 * @author tupog
 */
public class main {
    public static void main(String[] args) {
        Kendaraan[] kendaraanList = new Kendaraan[3];

        kendaraanList[0] = new Kendaraan("Kendaraan Umum");
        kendaraanList[1] = new Mobil("Toyota");
        kendaraanList[2] = new Motor("Yamaha");
        for (Kendaraan kd : kendaraanList) {
        System.out.println(kd.getInfo());
    }
  }
}
