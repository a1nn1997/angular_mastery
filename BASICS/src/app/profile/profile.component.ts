import { Component } from '@angular/core';

@Component({
  selector: 'test-1',
  templateUrl: './profile.component.html',
})
export class ProfileComponent {
  name:string = "Poppy";
  age:number = 25;
  ct:number=0;
  status:string = "coder";
  salary:number = 50000;
  getpug(){
    return 'I love you bebo';
  }
  but_prop:string ="red btn-large";
  btn_col1:string;
  conv_to_inr(){
   //alert(this.salary*80);
   this.ct+=1;
   if(this.ct%2==1){
    this.salary=this.salary*80;
   }
   else{
    this.salary=this.salary/80;
   }
  }
   inputValue:string="poppy"
   getInput(e:any){
    this.inputValue=e.target.value;
   }
   
   fruits:string[]=['aalu','apple','mango','grapes','banana']
  isDisable:boolean = true;
  my_date!:Date;
  my_date1:Date=new Date();
  constructor(){
    setTimeout(()=>{this.isDisable= false;},3000)
    const colors:Array<string> =["red","green","yellow","blue"]
   this.btn_col1=colors[Math.floor(Math.random()*4)]+" btn-large"
  }
  }
