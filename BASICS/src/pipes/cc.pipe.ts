import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'cc'
})
export class CcPipe implements PipeTransform {

  transform(value: string, code?:string): unknown {
    switch (code){
      case("USA"):{
        return"+1-"+value;
      }
      case("IND"):{
        return"+91-"+value;
      }
      case("PORK"):{
        return"+92-"+value;
      }
      default:{
        return value;
      }
    }
    
  }

}
