/* global google */
import React, { Component } from 'react';
import { withGoogleMap, GoogleMap, Marker, Circle } from 'react-google-maps';
import Button from '../../atom/Button';

import api from '../../../api';

import './style.css';

const initialZoomLevel = 13;
const kyotoLocation = { lat: 35.02107, lng: 135.75385 };

function markerData(eventInfo) {
  const title = eventInfo.title || 'no title';
  const position = { lat: Number(eventInfo.lat) || 0, lng: Number(eventInfo.lon) || 0 };
  const eventUrl = eventInfo.event_url
  return {
    title,
    position,
    eventUrl
  };
}

const GettingStartedGoogleMap = withGoogleMap((props) => {
  return (
    <GoogleMap
      ref={props.onMapLoad}
      defaultZoom={initialZoomLevel}
      defaultCenter={props.defaultCenter}
    >
      {
        props.eventInfo.map((eventInfo, i) => (
          <Marker
            { ...eventInfo }
            onClick={props.markerOnClick(eventInfo.eventUrl)}
            key={i}
          />
        ))
      }
      <Marker
        { ...props.userInfo }
        title="me"
      />
    </GoogleMap>
  );
});


export default class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      showEvent: false,
      defaultCenter: kyotoLocation,
      eventInfo: [],
      userInfo: {
        loading: true,
        position: {
          name: '',
          lat: 0,
          lng: 0
        },
      },
      markers: [],
    };
    this.handleMapLoad = this.handleMapLoad.bind(this);
    this.getDataSucceedHoc = this.getDataSucceedHoc.bind(this);
    this.showEvent = this.showEvent.bind(this);
  }

  componentWillMount() {
    this.getUserLocation()
      .then(this.getUserNearEvents)
      .then((res) => {
        const eventInfo = res.data.events;
        this.setState({ eventInfo })
      });
  }

  getUserLocation() {
    return new Promise((resolve, reject) => {
      navigator.geolocation.getCurrentPosition(
        this.getDataSucceedHoc(resolve),
        this.getDataFail
      );
    });
  }

  getDataSucceedHoc(next) {
    return (data) => {
      const { latitude, longitude } = data.coords;
      this.setState({
        userInfo: {
          position: {
            lat: latitude,
            lng: longitude
          },
          key: `My place`,
          defaultAnimation: 2,
        }
      }, next({ lat: latitude, lng: longitude }));
    }
  }

  getDataFail() {
    console.error('fail');
  }

  getUserNearEvents(location) {
    return api.getNearEvents(location)
      .then((res) => res);
  }

  handleMapLoad(map) {
    this._mapComponent = map;
    if (map) {
      // console.log(map.getZoom());
    }
  }

  showEvent() {
    this.setState({ showEvent: true });
  }

  markerOnClick(eventUrl) {
    return () => {
      window.open(eventUrl, "_blank")
    };
  }

  render() {
    const { userInfo, defaultCenter, showEvent } = this.state;
    const eventInfo = this.state.eventInfo.map(markerData);
    return (
      <div className="row home-container">
        <section className="col span-1-of-6"/>
        <section className="col span-2-of-3 home-main">
          {
            userInfo.loading ?
            <div className="sub-title">
              Now loading...
            </div> :
            <div className="sub-title">
              Find IT conferences near you!
              <Button
                onClick={this.showEvent}
                text="Find now!"
              />
            </div> 
          }
          <GettingStartedGoogleMap
            containerElement={
              <div className="googlemap-container" style={{ height: `100%` }} />
            }
            mapElement={
              <div style={{ height: `100%` }} />
            }
            onMapLoad={this.handleMapLoad}
            eventInfo={showEvent ? eventInfo : []}
            userInfo={userInfo}
            defaultCenter={defaultCenter}
            markerOnClick={this.markerOnClick}
          />
        </section>
        <section className="col span-1-of-6"/>
      </div>     
    );
  }
}