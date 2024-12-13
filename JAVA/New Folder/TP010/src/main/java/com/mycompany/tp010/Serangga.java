/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.tp010;

/**
 *
 * @author tupog
 */
public class Serangga {
    private final int jumlahKaki = 6;
    private String warna;

    public Serangga(String warna) {
        this.warna = warna;
    }

    public int getJumlahKaki() {
        return jumlahKaki;
    }

    public String getWarna() {
        return warna;
    }

    public void setWarna(String warna) {
        this.warna = warna;
    }

    public void gerak(int[] sumbu) {
        // Implementasi di kelas turunan
    }

    public void info() {
        // Implementasi di kelas turunan
    }
}
