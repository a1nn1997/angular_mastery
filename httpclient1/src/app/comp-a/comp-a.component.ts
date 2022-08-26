import { Component, OnInit } from '@angular/core';
import { CounterService } from 'src/services/counter.service';

@Component({
  selector: 'app-comp-a',
  templateUrl: './comp-a.component.html',
  styleUrls: ['./comp-a.component.css']
})
export class CompAComponent implements OnInit {
 
  
  constructor(private c:CounterService) { }
  
  ngOnInit(): void {
  }
  showCounter(){
    return this.c.getCounter()
  }
  updateCounter(){
    return this.c.updateCounter()
  }
}
