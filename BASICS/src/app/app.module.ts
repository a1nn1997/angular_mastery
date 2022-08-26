import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import {FormsModule} from '@angular/forms'
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

// import {ProfileComponent} from './profile/profile.component';
import { UsersComponent } from './users/users.component';
import { HighlightDirective } from '../directives/highlight.directive';
import { CcPipe } from '../pipes/cc.pipe'
@NgModule({
  declarations: [
    AppComponent,
    // ProfileComponent,
    UsersComponent,
    HighlightDirective,
    CcPipe
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
