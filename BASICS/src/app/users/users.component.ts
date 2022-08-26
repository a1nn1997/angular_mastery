import { Component, Input, Output, EventEmitter, OnInit, OnChanges, OnDestroy  } from '@angular/core';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent {//implements OnInit,OnChanges,OnDestroy {
  @Input() name!: string;
  @Input() reason!: string;
  @Input() img!:string;
  @Input() no!:number;

  @Output() myevent1 =new EventEmitter<string>();
  passData1(){
    this.myevent1.emit(`${this.name} FROM INPUT`)
    this.myevent1.emit(" FROM INPUT")
  }

  text!:string;
  //constructor was created/ updated at the time of instance was initialized, while ngOnInit at lifecycle  or component creation ie @Input @Output will not be visible to constructor but visible to ngOnInit container   constructor <  ngOnChanges < ngOnInit
  //ngOnDestroy when memory was freed
  constructor() {
   
  }
/*  listenerRef = setInterval(()=>{},1000);
  ngOnDestroy(){
      clearInterval(this.listenerRef)
  }

  ngOnChanges(){

  }
  ngOnInit(){
    this.text="starting"
   this.listenerRef = setInterval(()=>{
      console.log("timer.......")
    },3000)
    //constructor+extra
    //properties + event listener register+ initial data fetch   like component did mount
  }
  */  
}
