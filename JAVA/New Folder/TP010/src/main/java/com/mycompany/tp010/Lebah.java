/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.tp010;

/**
 *
 * @author tupog
 */
public class Lebah extends Serangga implements BisaTerbang {
    private Koordinat3D posisi;

    public Lebah(String warna, int x, int y, int z) {
        super(warna);
        this.posisi = new Koordinat3D(x, y, z);
    }

    @Override
    public void gerak(int[] sumbu) {
        posisi.setX(sumbu[0]);
        posisi.setY(sumbu[1]);
        posisi.setZ(sumbu[2]);
    }

    @Override
    public void terbang(int x, int y, int z) {
        posisi.setX(posisi.getX() + x);
        posisi.setY(posisi.getY() + y);
        posisi.setZ(posisi.getZ() + z);
    }

    @Override
    public void info() {
        System.out.println("Warna Lebah : " + getWarna());
        System.out.println("Posisi Lebah : X = " + posisi.getX() + ", Y = " + posisi.getY() + ", Z = " + posisi.getZ());
    }
}

