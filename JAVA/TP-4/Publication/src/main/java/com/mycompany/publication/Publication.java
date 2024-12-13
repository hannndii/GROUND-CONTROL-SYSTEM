/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.publication;
import java.text.NumberFormat;
/*Kelas java.text dalam Java menyediakan berbagai kelas dan metode untuk pemformatan dan analisis teks, 
termasuk pemformatan angka, mata uang, tanggal, dan waktu. Kelas-kelas dalam paket ini membantu dalam mengubah 
representasi data dalam bentuk string agar sesuai dengan format lokal atau kebutuhan pengguna. */

import java.util.Locale;
/**
 *
 * @author tupog
 */
public class Publication {
    private String title;
    private int price;
    private int page;
    
    // Locale Indonesia untuk format mata uang (Rp)
    Locale myIndonesianLocal = new Locale("in", "ID");
    NumberFormat formatter = NumberFormat.getCurrencyInstance(myIndonesianLocal);
    
    void setTitle(String title){
        this.title = title;
    }
    
    String getTitle(){
        return title;
    }
    
    void setPrice(int price){
        this.price = price;
    }
    
    int getPrice(){
        return price;
    }
    
    void setPage(int page){
        this.page = page;
    }
    
    int getPage(){
        return page;
    }
    
    void printHeader(){
        System.out.println("=====Mencari Informasi Buku======\n");
    }
    
    void printBookPrice(){
        System.out.println(getTitle() + String.format("(Price List: %s)", formatter.format(getPrice())));
    }
    
    void printBookPage(){
        System.out.println(getTitle() + " - Pembelian Jilid: " + getPage());
    }
            
    void printBookPurchase(){
        int result = page * price;
        System.out.println(getTitle() + " " + getPage() + " Jilid" + String.format("(%s)", formatter.format(result)));
    } 
}
