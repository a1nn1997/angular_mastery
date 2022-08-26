import { Directive, ElementRef, HostBinding, HostListener, OnInit } from '@angular/core';

@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective {
  constructor(private el: ElementRef) { 
    //private el is constructor injection where initializing and this.el=el was done in short
   }
   
   @HostBinding('style.backgroundColor') bgColor:any;  //@hostbinding as ngColor as variable
   @HostListener('mouseenter') onEnter(){
    this.bgColor='blue'
   }
   @HostListener('mouseleave') onLeave(){
    this.bgColor='pink'
   }

   ngOnInit(){
    //this.el.nativeElement.style.backgroundColor= "pink"
    //this.bgColor='pink'
   }
}
