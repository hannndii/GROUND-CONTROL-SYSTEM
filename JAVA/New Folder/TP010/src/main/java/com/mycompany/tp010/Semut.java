/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.tp010;

/**
 *
 * @author tupog
 */
public class Semut extends Serangga {
    private Koordinat posisi;

    public Semut(String warna, int x, int y) {
        super(warna);
        this.posisi = new Koordinat(x, y);
    }

    @Override
    public void gerak(int[] sumbu) {
        posisi.setX(sumbu[0]);
        posisi.setY(sumbu[1]);
    }

    @Override
    public void info() {
        System.out.println("Warna Semut : " + getWarna());
        System.out.println("Posisi Semut : X = " + posisi.getX() + ", Y = " + posisi.getY());
    }
}