/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.product;
import java.text.NumberFormat;
import java.util.Locale;
/**
 *
 * @author tupog
 */
public class Product {
    private String namaProduk;
    private double harga;
    private int stok;
    
    Locale myIndonesianLocale = new Locale("id", "ID");
    NumberFormat formatter = NumberFormat.getCurrencyInstance(myIndonesianLocale);
    
    public Product(String namaProduk, double harga, int stok){
        this.namaProduk = namaProduk;
        this.harga = harga;
        this.stok = stok;
    }
    
    void setNameProduct(String namaProduk){
        this.namaProduk = namaProduk;
    }
    
    String getNameProduct(){
        return namaProduk;
    }
    
    void setHarga(double harga){
        this.harga = harga;
    }
    
    double getHarga(){
        return harga;
    }
    
    void setStock(int stok){
        this.stok = stok;
    }
    
    int getStock(){
        return stok;
    }
    
    void printProductInfo(){
        System.out.println("Nama Produk: " + getNameProduct());
        System.out.println("Harga: " + formatter.format(getHarga()));    
        System.out.println("Stok: " + getStock() + " unit");
    }
}
