import { Component } from '@angular/core';  
import {User_temp1} from 'src/models/user'  //absolute path   '../models/user'    //relative path
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  phoneNo:string="1234567898";
  title:string = 'poppy';
  user_comp1: Array<User_temp1>= [
    {name:"poppy",reason:"yaha hm data parents se child ko bhej rhe he",imagePath:"assets/1.jpeg",no:1},
    {name:"pogli aurat",reason:"just like state in react",imagePath:"assets/2.jpeg",no:2},
    {name:"bebo",reason:"with for loop",imagePath:"assets/3.jpg",no:3},
  ]
  received_comp1(e:any){
    console.log(e)
  }}
