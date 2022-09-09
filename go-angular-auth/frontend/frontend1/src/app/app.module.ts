import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';

import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/user/login/login.component';
import { RegisterComponent } from './components/user/register/register.component';

import { HttpClientModule} from '@angular/common/http';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { NgxWebstorageModule } from 'ngx-webstorage';
import { ToastrModule } from 'ngx-toastr';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { EditorModule } from '@tinymce/tinymce-angular'
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { EdituserComponent } from './components/user/edituser/edituser.component';
import { CreatePostComponent } from './components/post/create-post/create-post.component';
import { ViewPostComponent } from './components/post/view-post/view-post.component';
import { PostTitleComponent } from './components/shared/post-title/post-title.component';
import { SubredditSideBarComponent } from './components/shared/subreddit-side-bar/subreddit-side-bar.component';
import { SideBarComponent } from './components/shared/side-bar/side-bar.component';
import { VoteButtonComponent } from './components/shared/vote-button/vote-button.component';
import { CreateSubredditComponent } from './components/subreddit/create-subreddit/create-subreddit.component';
import { ListSubredditComponent } from './components/subreddit/list-subreddit/list-subreddit.component';
import { UserProfileComponent } from './components/user/user-profile/user-profile.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    HomeComponent,
    LoginComponent,
    RegisterComponent,
    EdituserComponent,
    CreatePostComponent,
    ViewPostComponent,
    PostTitleComponent,
    SubredditSideBarComponent,
    SideBarComponent,
    VoteButtonComponent,
    CreateSubredditComponent,
    ListSubredditComponent,
    UserProfileComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    NgxWebstorageModule.forRoot(),
    BrowserAnimationsModule,
    ToastrModule.forRoot(),
    FontAwesomeModule,
    EditorModule,
    NgbModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
