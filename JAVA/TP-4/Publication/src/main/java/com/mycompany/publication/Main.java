/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.publication;
import java.text.NumberFormat;
import java.util.Locale;

/**
 *
 * @author tupog
 */
public class Main {

    public static void main(String[] args) {
        Locale myIndonesianLocal = new Locale("in", "ID");
        NumberFormat formater = NumberFormat.getCurrencyInstance(myIndonesianLocal);
        
        Publication dataBook1 = new Publication();
        dataBook1.setTitle("Three Kingdom");
        dataBook1.setPrice(5000);
        dataBook1.printHeader();
        dataBook1.printBookPrice();
        
        Publication dataBook2 = new Publication();
        dataBook2.setTitle("a Pie");
        dataBook2.setPrice(8400);
        dataBook2.printBookPrice();
        System.out.println("\n=====================\n");
        
        dataBook1.getTitle();
        dataBook1.setPage(67);
        dataBook1.printBookPage();
        dataBook2.getTitle();
        dataBook2.setPage(82);
        dataBook2.printBookPage();
        System.out.println("\n=====================\n");
        
        dataBook1.getPage();
        dataBook1.printBookPurchase();
         dataBook2.getPage();
        dataBook2.printBookPurchase();
        System.out.println("\n=====================\n");
        
        int total = (dataBook1.getPrice()*dataBook1.getPage()) + (dataBook2.getPrice() * dataBook2.getPage());
        System.out.println("Jumlah total pembelian: "+ formater.format(total));       
    }
}
