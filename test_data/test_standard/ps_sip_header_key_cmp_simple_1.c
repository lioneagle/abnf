#include "ps_sip_header_key_cmp_simple_1.h"

unsigned int GetSipHeaderIndex(char const** src, char const* end)
{
	char const*  p = *src;
	
	if (p == NULL || p >= end) {
		return ABNF_SIP_HDR_UNKNOWN;
	}

	switch (*(p++) | 0x20) {
		case 'a':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_ACCEPT_CONTACT;
			}
			switch (*(p++) | 0x20) {
				case 'c':
					if ((p < end) && ((*(p++) | 0x20) == 'c') {
						if ((p < end) && ((*(p++) | 0x20) == 'e') {
							if ((p < end) && ((*(p++) | 0x20) == 'p') {
								if ((p < end) && ((*(p++) | 0x20) == 't') {
									if ((p < end) && (*(p++) == '-') {
										if ((p < end) && ((*(p++) | 0x20) == 'c') {
											if ((p < end) && ((*(p++) | 0x20) == 'o') {
												if ((p < end) && ((*(p++) | 0x20) == 'n') {
													if ((p < end) && ((*(p++) | 0x20) == 't') {
														if ((p < end) && ((*(p++) | 0x20) == 'a') {
															if ((p < end) && ((*(p++) | 0x20) == 'c') {
																if ((p < end) && ((*(p++) | 0x20) == 't') {
																	if (p >= end) {
																		*src = p;
																		return ABNF_SIP_HDR_ACCEPT_CONTACT;
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
				case 'l':
					if ((p < end) && ((*(p++) | 0x20) == 'l') {
						if ((p < end) && ((*(p++) | 0x20) == 'o') {
							if ((p < end) && ((*(p++) | 0x20) == 'w') {
								if (p >= end) {
									*src = p;
									return ABNF_SIP_HDR_ALLOW;
								}
								if ((p < end) && (*(p++) == '-') {
									if ((p < end) && ((*(p++) | 0x20) == 'e') {
										if ((p < end) && ((*(p++) | 0x20) == 'v') {
											if ((p < end) && ((*(p++) | 0x20) == 'e') {
												if ((p < end) && ((*(p++) | 0x20) == 'n') {
													if ((p < end) && ((*(p++) | 0x20) == 't') {
														if ((p < end) && ((*(p++) | 0x20) == 's') {
															if (p >= end) {
																*src = p;
																return ABNF_SIP_HDR_ALLOW_EVENTS;
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'b':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_REFERRED_BY;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'c':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_CONTENT_TYPE;
			}
			switch (*(p++) | 0x20) {
				case 'a':
					if ((p < end) && ((*(p++) | 0x20) == 'l') {
						if ((p < end) && ((*(p++) | 0x20) == 'l') {
							if ((p < end) && (*(p++) == '-') {
								if ((p < end) && ((*(p++) | 0x20) == 'i') {
									if ((p < end) && ((*(p++) | 0x20) == 'd') {
										if (p >= end) {
											*src = p;
											return ABNF_SIP_HDR_CALL_ID;
										}
									}
								}
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
				case 'o':
					if ((p < end) && ((*(p++) | 0x20) == 'n') {
						if ((p < end) && ((*(p++) | 0x20) == 't') {
							switch (*(p++) | 0x20) {
								case 'a':
									if ((p < end) && ((*(p++) | 0x20) == 'c') {
										if ((p < end) && ((*(p++) | 0x20) == 't') {
											if (p >= end) {
												*src = p;
												return ABNF_SIP_HDR_CONTACT;
											}
										}
									}
									*src = p;
									return ABNF_SIP_HDR_UNKNOWN;
								case 'e':
									if ((p < end) && ((*(p++) | 0x20) == 'n') {
										if ((p < end) && ((*(p++) | 0x20) == 't') {
											if ((p < end) && (*(p++) == '-') {
												switch (*(p++) | 0x20) {
													case 'd':
														if ((p < end) && ((*(p++) | 0x20) == 'i') {
															if ((p < end) && ((*(p++) | 0x20) == 's') {
																if ((p < end) && ((*(p++) | 0x20) == 'p') {
																	if ((p < end) && ((*(p++) | 0x20) == 'o') {
																		if ((p < end) && ((*(p++) | 0x20) == 's') {
																			if ((p < end) && ((*(p++) | 0x20) == 'i') {
																				if ((p < end) && ((*(p++) | 0x20) == 't') {
																					if ((p < end) && ((*(p++) | 0x20) == 'i') {
																						if ((p < end) && ((*(p++) | 0x20) == 'o') {
																							if ((p < end) && ((*(p++) | 0x20) == 'n') {
																								if (p >= end) {
																									*src = p;
																									return ABNF_SIP_HDR_CONTENT_DISPOSITION;
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
														*src = p;
														return ABNF_SIP_HDR_UNKNOWN;
													case 'e':
														if ((p < end) && ((*(p++) | 0x20) == 'n') {
															if ((p < end) && ((*(p++) | 0x20) == 'c') {
																if ((p < end) && ((*(p++) | 0x20) == 'o') {
																	if ((p < end) && ((*(p++) | 0x20) == 'd') {
																		if ((p < end) && ((*(p++) | 0x20) == 'i') {
																			if ((p < end) && ((*(p++) | 0x20) == 'n') {
																				if ((p < end) && ((*(p++) | 0x20) == 'g') {
																					if (p >= end) {
																						*src = p;
																						return ABNF_SIP_HDR_CONTENT_ENCODING;
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
														*src = p;
														return ABNF_SIP_HDR_UNKNOWN;
													case 'l':
														if ((p < end) && ((*(p++) | 0x20) == 'e') {
															if ((p < end) && ((*(p++) | 0x20) == 'n') {
																if ((p < end) && ((*(p++) | 0x20) == 'g') {
																	if ((p < end) && ((*(p++) | 0x20) == 't') {
																		if ((p < end) && ((*(p++) | 0x20) == 'h') {
																			if (p >= end) {
																				*src = p;
																				return ABNF_SIP_HDR_CONTENT_LENGTH;
																			}
																		}
																	}
																}
															}
														}
														*src = p;
														return ABNF_SIP_HDR_UNKNOWN;
													case 't':
														if ((p < end) && ((*(p++) | 0x20) == 'y') {
															if ((p < end) && ((*(p++) | 0x20) == 'p') {
																if ((p < end) && ((*(p++) | 0x20) == 'e') {
																	if (p >= end) {
																		*src = p;
																		return ABNF_SIP_HDR_CONTENT_TYPE;
																	}
																}
															}
														}
														*src = p;
														return ABNF_SIP_HDR_UNKNOWN;
												}
											}
										}
									}
									*src = p;
									return ABNF_SIP_HDR_UNKNOWN;
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
				case 's':
					if ((p < end) && ((*(p++) | 0x20) == 'e') {
						if ((p < end) && ((*(p++) | 0x20) == 'q') {
							if (p >= end) {
								*src = p;
								return ABNF_SIP_HDR_CSEQ;
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'd':
			if ((p < end) && ((*(p++) | 0x20) == 'a') {
				if ((p < end) && ((*(p++) | 0x20) == 't') {
					if ((p < end) && ((*(p++) | 0x20) == 'e') {
						if (p >= end) {
							*src = p;
							return ABNF_SIP_HDR_DATE;
						}
					}
				}
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'e':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_CONTENT_ENCODING;
			}
			if ((p < end) && ((*(p++) | 0x20) == 'v') {
				if ((p < end) && ((*(p++) | 0x20) == 'e') {
					if ((p < end) && ((*(p++) | 0x20) == 'n') {
						if ((p < end) && ((*(p++) | 0x20) == 't') {
							if (p >= end) {
								*src = p;
								return ABNF_SIP_HDR_EVENT;
							}
						}
					}
				}
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'f':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_FROM;
			}
			if ((p < end) && ((*(p++) | 0x20) == 'r') {
				if ((p < end) && ((*(p++) | 0x20) == 'o') {
					if ((p < end) && ((*(p++) | 0x20) == 'm') {
						if (p >= end) {
							*src = p;
							return ABNF_SIP_HDR_FROM;
						}
					}
				}
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'i':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_CALL_ID;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'j':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_REJECT_CONTACT;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'k':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_SUPPORTED;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'l':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_CONTENT_LENGTH;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'm':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_CONTACT;
			}
			switch (*(p++) | 0x20) {
				case 'a':
					if ((p < end) && ((*(p++) | 0x20) == 'x') {
						if ((p < end) && (*(p++) == '-') {
							if ((p < end) && ((*(p++) | 0x20) == 'f') {
								if ((p < end) && ((*(p++) | 0x20) == 'o') {
									if ((p < end) && ((*(p++) | 0x20) == 'r') {
										if ((p < end) && ((*(p++) | 0x20) == 'w') {
											if ((p < end) && ((*(p++) | 0x20) == 'a') {
												if ((p < end) && ((*(p++) | 0x20) == 'r') {
													if ((p < end) && ((*(p++) | 0x20) == 'd') {
														if ((p < end) && ((*(p++) | 0x20) == 's') {
															if (p >= end) {
																*src = p;
																return ABNF_SIP_HDR_MAX_FORWARDS;
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
				case 'i':
					if ((p < end) && ((*(p++) | 0x20) == 'm') {
						if ((p < end) && ((*(p++) | 0x20) == 'e') {
							if ((p < end) && (*(p++) == '-') {
								if ((p < end) && ((*(p++) | 0x20) == 'v') {
									if ((p < end) && ((*(p++) | 0x20) == 'e') {
										if ((p < end) && ((*(p++) | 0x20) == 'r') {
											if ((p < end) && ((*(p++) | 0x20) == 's') {
												if ((p < end) && ((*(p++) | 0x20) == 'i') {
													if ((p < end) && ((*(p++) | 0x20) == 'o') {
														if ((p < end) && ((*(p++) | 0x20) == 'n') {
															if (p >= end) {
																*src = p;
																return ABNF_SIP_HDR_MIME_VERSION;
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'o':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_EVENT;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'r':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_REFER_TO;
			}
			switch (*(p++) | 0x20) {
				case 'e':
					switch (*(p++) | 0x20) {
						case 'c':
							if ((p < end) && ((*(p++) | 0x20) == 'o') {
								if ((p < end) && ((*(p++) | 0x20) == 'r') {
									if ((p < end) && ((*(p++) | 0x20) == 'd') {
										if ((p < end) && (*(p++) == '-') {
											if ((p < end) && ((*(p++) | 0x20) == 'r') {
												if ((p < end) && ((*(p++) | 0x20) == 'o') {
													if ((p < end) && ((*(p++) | 0x20) == 'u') {
														if ((p < end) && ((*(p++) | 0x20) == 't') {
															if ((p < end) && ((*(p++) | 0x20) == 'e') {
																if (p >= end) {
																	*src = p;
																	return ABNF_SIP_HDR_RECORD_ROUTE;
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
							*src = p;
							return ABNF_SIP_HDR_UNKNOWN;
						case 'f':
							if ((p < end) && ((*(p++) | 0x20) == 'e') {
								if ((p < end) && ((*(p++) | 0x20) == 'r') {
									switch (*(p++) | 0x20) {
										case '-':
											if ((p < end) && ((*(p++) | 0x20) == 't') {
												if ((p < end) && ((*(p++) | 0x20) == 'o') {
													if (p >= end) {
														*src = p;
														return ABNF_SIP_HDR_REFER_TO;
													}
												}
											}
											*src = p;
											return ABNF_SIP_HDR_UNKNOWN;
										case 'r':
											if ((p < end) && ((*(p++) | 0x20) == 'e') {
												if ((p < end) && ((*(p++) | 0x20) == 'd') {
													if ((p < end) && (*(p++) == '-') {
														if ((p < end) && ((*(p++) | 0x20) == 'b') {
															if ((p < end) && ((*(p++) | 0x20) == 'y') {
																if (p >= end) {
																	*src = p;
																	return ABNF_SIP_HDR_REFERRED_BY;
																}
															}
														}
													}
												}
											}
											*src = p;
											return ABNF_SIP_HDR_UNKNOWN;
									}
								}
							}
							*src = p;
							return ABNF_SIP_HDR_UNKNOWN;
						case 'j':
							if ((p < end) && ((*(p++) | 0x20) == 'e') {
								if ((p < end) && ((*(p++) | 0x20) == 'c') {
									if ((p < end) && ((*(p++) | 0x20) == 't') {
										if ((p < end) && (*(p++) == '-') {
											if ((p < end) && ((*(p++) | 0x20) == 'c') {
												if ((p < end) && ((*(p++) | 0x20) == 'o') {
													if ((p < end) && ((*(p++) | 0x20) == 'n') {
														if ((p < end) && ((*(p++) | 0x20) == 't') {
															if ((p < end) && ((*(p++) | 0x20) == 'a') {
																if ((p < end) && ((*(p++) | 0x20) == 'c') {
																	if ((p < end) && ((*(p++) | 0x20) == 't') {
																		if (p >= end) {
																			*src = p;
																			return ABNF_SIP_HDR_REJECT_CONTACT;
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
							*src = p;
							return ABNF_SIP_HDR_UNKNOWN;
						case 'q':
							if ((p < end) && ((*(p++) | 0x20) == 'u') {
								if ((p < end) && ((*(p++) | 0x20) == 'e') {
									if ((p < end) && ((*(p++) | 0x20) == 's') {
										if ((p < end) && ((*(p++) | 0x20) == 't') {
											if ((p < end) && (*(p++) == '-') {
												if ((p < end) && ((*(p++) | 0x20) == 'd') {
													if ((p < end) && ((*(p++) | 0x20) == 'i') {
														if ((p < end) && ((*(p++) | 0x20) == 's') {
															if ((p < end) && ((*(p++) | 0x20) == 'p') {
																if ((p < end) && ((*(p++) | 0x20) == 'o') {
																	if ((p < end) && ((*(p++) | 0x20) == 's') {
																		if ((p < end) && ((*(p++) | 0x20) == 'i') {
																			if ((p < end) && ((*(p++) | 0x20) == 't') {
																				if ((p < end) && ((*(p++) | 0x20) == 'i') {
																					if ((p < end) && ((*(p++) | 0x20) == 'o') {
																						if ((p < end) && ((*(p++) | 0x20) == 'n') {
																							if (p >= end) {
																								*src = p;
																								return ABNF_SIP_HDR_REQUEST_DISPOSITION;
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
							*src = p;
							return ABNF_SIP_HDR_UNKNOWN;
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
				case 'o':
					if ((p < end) && ((*(p++) | 0x20) == 'u') {
						if ((p < end) && ((*(p++) | 0x20) == 't') {
							if ((p < end) && ((*(p++) | 0x20) == 'e') {
								if (p >= end) {
									*src = p;
									return ABNF_SIP_HDR_ROUTE;
								}
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 's':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_SUBJECT;
			}
			switch (*(p++) | 0x20) {
				case 'e':
					if ((p < end) && ((*(p++) | 0x20) == 's') {
						if ((p < end) && ((*(p++) | 0x20) == 's') {
							if ((p < end) && ((*(p++) | 0x20) == 'i') {
								if ((p < end) && ((*(p++) | 0x20) == 'o') {
									if ((p < end) && ((*(p++) | 0x20) == 'n') {
										if ((p < end) && (*(p++) == '-') {
											if ((p < end) && ((*(p++) | 0x20) == 'e') {
												if ((p < end) && ((*(p++) | 0x20) == 'x') {
													if ((p < end) && ((*(p++) | 0x20) == 'p') {
														if ((p < end) && ((*(p++) | 0x20) == 'i') {
															if ((p < end) && ((*(p++) | 0x20) == 'r') {
																if ((p < end) && ((*(p++) | 0x20) == 'e') {
																	if ((p < end) && ((*(p++) | 0x20) == 's') {
																		if (p >= end) {
																			*src = p;
																			return ABNF_SIP_HDR_SESSION_EXPIRES;
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
				case 'u':
					switch (*(p++) | 0x20) {
						case 'b':
							if ((p < end) && ((*(p++) | 0x20) == 'j') {
								if ((p < end) && ((*(p++) | 0x20) == 'e') {
									if ((p < end) && ((*(p++) | 0x20) == 'c') {
										if ((p < end) && ((*(p++) | 0x20) == 't') {
											if (p >= end) {
												*src = p;
												return ABNF_SIP_HDR_SUBJECT;
											}
										}
									}
								}
							}
							*src = p;
							return ABNF_SIP_HDR_UNKNOWN;
						case 'p':
							if ((p < end) && ((*(p++) | 0x20) == 'p') {
								if ((p < end) && ((*(p++) | 0x20) == 'o') {
									if ((p < end) && ((*(p++) | 0x20) == 'r') {
										if ((p < end) && ((*(p++) | 0x20) == 't') {
											if ((p < end) && ((*(p++) | 0x20) == 'e') {
												if ((p < end) && ((*(p++) | 0x20) == 'd') {
													if (p >= end) {
														*src = p;
														return ABNF_SIP_HDR_SUPPORTED;
													}
												}
											}
										}
									}
								}
							}
							*src = p;
							return ABNF_SIP_HDR_UNKNOWN;
					}
					*src = p;
					return ABNF_SIP_HDR_UNKNOWN;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 't':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_TO;
			}
			if ((p < end) && ((*(p++) | 0x20) == 'o') {
				if (p >= end) {
					*src = p;
					return ABNF_SIP_HDR_TO;
				}
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'u':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_ALLOW_EVENTS;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'v':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_VIA;
			}
			if ((p < end) && ((*(p++) | 0x20) == 'i') {
				if ((p < end) && ((*(p++) | 0x20) == 'a') {
					if (p >= end) {
						*src = p;
						return ABNF_SIP_HDR_VIA;
					}
				}
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
		case 'x':
			if (p >= end) {
				*src = p;
				return ABNF_SIP_HDR_SESSION_EXPIRES;
			}
			*src = p;
			return ABNF_SIP_HDR_UNKNOWN;
	}

	return ABNF_SIP_HDR_UNKNOWN;
}
